# bitcalc — Bitwise / Arithmetic Expression Calculator

**Date:** 2026-06-25
**Status:** Approved (design)

## Purpose

Add a third CLI tool, `bitcalc`, to the hexdec-converter project for evaluating
integer expressions — primarily bit operations like `1 << 21`, `1 & 0x03`, and
`1 << 2 & 0x03` — and printing the result in decimal, hex, and binary at once.

## Architecture

Mirrors the existing layout (core logic in `pkg/`, thin CLI in `cmd/`):

- **`pkg/calculator/calculator.go`** — self-contained evaluator exposing a single
  public function `Eval(expr string) (uint64, error)`. Internally a tokenizer feeds
  a recursive-descent parser. Stdlib only. Independently unit-testable.
- **`cmd/bitcalc/main.go`** — thin wrapper. Joins `os.Args[1:]` into one expression
  (so `bitcalc 1 + 2` works unquoted; shell metacharacters like `<<`, `&`, `|`, `(`
  still require quoting by the user), calls `Eval`, prints three lines.

## Numeric Model

All evaluation is **`uint64` with wraparound** (Go's native unsigned semantics),
consistent with the existing `converter` package's `uint64` world.

- `~0` → `18446744073709551615` (`0xFFFFFFFFFFFFFFFF`)
- `1 << 21` → `2097152`
- `1 - 2` wraps (standard fixed-width programmer-calculator behavior)
- Division or modulo by zero → error

## Number Literals

- Decimal: `42`
- Hex: `0x03` / `0X03`
- Binary: `0b1010` / `0B1010`

## Operators & Precedence (C-style, high → low)

| Level          | Operators        |
|----------------|------------------|
| unary          | `~`  `-`  `+`    |
| multiplicative | `*`  `/`  `%`    |
| additive       | `+`  `-`         |
| shift          | `<<`  `>>`       |
| bitwise AND    | `&`              |
| bitwise XOR    | `^`              |
| bitwise OR     | `\|`             |

Parentheses override precedence. Worked example: `1 << 2 & 0x03` parses as
`(1 << 2) & 0x03` = `4 & 3` = `0`.

## Output

Three lines — decimal, hex (`0x` prefix), binary (`0b` prefix):

```
$ bitcalc '1 << 21'
2097152
0x200000
0b1000000000000000000000
```

## Error Handling

Invalid tokens, unbalanced parentheses, missing/trailing operands, and div/mod by
zero print `Error: <message>` and exit with status 1, matching the existing
h2d/d2h convention. (Errors are written to stderr — a small, deliberate
improvement over h2d/d2h, which write them to stdout — so output piping stays
clean.)

## Testing

- `pkg/calculator/calculator_test.go` — table-driven: each literal form, each
  operator, precedence chains (incl. `1 << 2 & 0x03`), parentheses, and error cases.
- Add a `bitcalc` case to `test/intergration_test.go`.

## Project Integration

- Add `bitcalc` to `Makefile` `BINARIES`, the `install` targets, and `info` text.
- Document the new tool in `README.md` (usage, examples, project structure).

## Out of Scope (YAGNI)

- Floating-point math.
- Variables, functions, or REPL/interactive mode.
- Configurable bit widths (fixed at 64-bit).
- Signed-output formatting.
