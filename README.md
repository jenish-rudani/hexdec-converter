# Hexadecimal-Decimal Converter

This Go module provides command-line utilities for converting between hexadecimal and decimal numbers, plus `bitcalc` for evaluating bitwise and arithmetic expressions.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Examples](#examples)
- [Building](#building)
- [Testing](#testing)
- [Project Structure](#project-structure)
- [Development](#development)
  - [Makefile Commands](#makefile-commands)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/jenish-rudani/hexdec-converter.git
cd hexdec-converter
```

2. Build the binaries:

```bash
make
```

3. Install the utilities to your system PATH:

```bash
make install
```

Note: The `make install` command may require sudo privileges on Linux and macOS. On Windows, you may need to run your terminal as Administrator.

## Usage

After installation, you can use the utilities from anywhere in your terminal:

Convert hexadecimal to decimal:

```bash
h2d <hexadecimal number>
```

Convert decimal to hexadecimal:

```bash
d2h <decimal number>
```

Evaluate a bitwise/arithmetic expression:

```bash
bitcalc <expression>
```

`bitcalc` accepts decimal (`42`), hex (`0x2A`), and binary (`0b101010`) literals.
It supports the operators `~ << >> & | ^ + - * / %` and parentheses, following
C-style precedence, and all math is done in 64-bit unsigned arithmetic. The
result is printed in decimal, hex, and binary.

> Note: shell metacharacters such as `<<`, `&`, `|`, and `(` must be quoted, so
> wrap the expression in quotes: `bitcalc '1 << 2 & 0x03'`.

### Examples

Here are some examples of how to use the utilities:

1. Hexadecimal to Decimal:

   ```bash
   $ h2d A
   10

   $ h2d FF
   255

   $ h2d 100
   256
   ```

2. Decimal to Hexadecimal:

   ```bash
   $ d2h 10
   A

   $ d2h 255
   FF

   $ d2h 256
   100
   ```

3. Bitwise/Arithmetic Calculator:

   ```bash
   $ bitcalc '1 << 21'
   2097152
   0x200000
   0b1000000000000000000000

   $ bitcalc '1 & 0x03'
   1
   0x1
   0b1

   $ bitcalc '1 << 2 & 0x03'
   0
   0x0
   0b0

   $ bitcalc '~0'
   18446744073709551615
   0xFFFFFFFFFFFFFFFF
   0b1111111111111111111111111111111111111111111111111111111111111111
   ```

## Building

To build binaries for multiple platforms:

```bash
make
```

This will create binaries in the `binary` directory, organized by platform.

## Testing

To run the unit and integration tests:

```bash
make test
```

This will run tests for the converter package and integration tests for the command-line tools.

## Project Structure

```
hexdec-converter/
├── cmd/
│   ├── h2d/
│   │   └── main.go
│   ├── d2h/
│   │   └── main.go
│   └── bitcalc/
│       └── main.go
├── pkg/
│   ├── converter/
│   │   ├── converter.go
│   │   └── converter_test.go
│   └── calculator/
│       ├── calculator.go
│       └── calculator_test.go
├── test/
│   └── integration_test.go
├── Makefile
├── go.mod
└── README.md
```

- `cmd/`: Contains the main applications for h2d, d2h, and bitcalc.
- `pkg/`: Contains the core logic — the `converter` package (hex/dec conversion) and the `calculator` package (expression evaluation).
- `test/`: Contains integration tests for the command-line tools.

## Development

### Makefile Commands

The project includes a Makefile with the following commands:

- `make`: Build the binaries for all supported platforms
- `make test`: Run unit and integration tests
- `make clean`: Remove all generated binaries
- `make info`: Display usage information and examples
- `make install`: Install the utilities to your system PATH

To view the usage information and examples in the terminal, run:

```bash
make info
```

This will display:

```
Hexadecimal-Decimal Converter
Usage:
  h2d <hexadecimal number>  - Convert hex to decimal
  d2h <decimal number>      - Convert decimal to hex
  bitcalc <expression>      - Evaluate a bitwise/arithmetic expression

Examples:
  h2d FF             -> 255
  d2h 255            -> FF
  bitcalc '1 << 21'  -> 2097152 / 0x200000 / 0b...0

Build:   make
Test:    make test
Clean:   make clean
Install: make install
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[GNU General Public License v3.0](LICENSE)
```
See License file for more details
```