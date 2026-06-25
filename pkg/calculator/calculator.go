// ============================================================================
//
// # Hexdec Converter
//
// File:        calculator.go
// Description: Bitwise/arithmetic expression evaluator (recursive descent)
//
// Author:      Jenish Rudani
// Created:     2026-06-25
// Updated:     2026-06-25
//
// Version:     1.0
// Copyright:   (c) 2026 Jenish Rudani
// License:     GNU General Public License v3.0
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//
// ============================================================================

// Package calculator evaluates integer expressions made of bitwise and
// arithmetic operators. All math is performed in uint64 with wraparound, so
// the tool behaves like a fixed-width programmer's calculator (e.g. ~0 yields
// 0xFFFFFFFFFFFFFFFF). Operator precedence follows the C convention.
package calculator

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Eval parses and evaluates a single expression string, returning the result
// as a uint64. It reports an error for malformed input or division by zero.
func Eval(expr string) (uint64, error) {
	tokens, err := tokenize(expr)
	if err != nil {
		return 0, err
	}
	p := &parser{tokens: tokens}
	value, err := p.parseExpr()
	if err != nil {
		return 0, err
	}
	if !p.atEnd() {
		return 0, fmt.Errorf("unexpected token %q", p.peek().text)
	}
	return value, nil
}

// ---------------------------------------------------------------------------
// Tokenizer
// ---------------------------------------------------------------------------

type tokenKind int

const (
	tokNumber tokenKind = iota
	tokOp
	tokLParen
	tokRParen
)

type token struct {
	kind  tokenKind
	text  string
	value uint64 // populated for tokNumber
}

// tokenize splits an expression into tokens, recognising decimal, hex (0x) and
// binary (0b) literals plus single- and multi-character operators.
func tokenize(expr string) ([]token, error) {
	var tokens []token
	runes := []rune(expr)

	for i := 0; i < len(runes); {
		c := runes[i]

		switch {
		case unicode.IsSpace(c):
			i++

		case c == '(':
			tokens = append(tokens, token{kind: tokLParen, text: "("})
			i++

		case c == ')':
			tokens = append(tokens, token{kind: tokRParen, text: ")"})
			i++

		case c == '<' || c == '>':
			// Shift operators must come in pairs (<< or >>).
			if i+1 >= len(runes) || runes[i+1] != c {
				return nil, fmt.Errorf("unexpected character %q", string(c))
			}
			tokens = append(tokens, token{kind: tokOp, text: string([]rune{c, c})})
			i += 2

		case strings.ContainsRune("&|^~+-*/%", c):
			tokens = append(tokens, token{kind: tokOp, text: string(c)})
			i++

		case unicode.IsDigit(c):
			start := i
			for i < len(runes) && isNumberRune(runes[i]) {
				i++
			}
			literal := string(runes[start:i])
			value, err := parseNumber(literal)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token{kind: tokNumber, text: literal, value: value})

		default:
			return nil, fmt.Errorf("unexpected character %q", string(c))
		}
	}

	return tokens, nil
}

// isNumberRune reports whether r can appear inside a numeric literal (digits,
// hex letters, and the x/b base markers).
func isNumberRune(r rune) bool {
	if unicode.IsDigit(r) {
		return true
	}
	switch unicode.ToLower(r) {
	case 'a', 'b', 'c', 'd', 'e', 'f', 'x':
		return true
	}
	return false
}

// parseNumber converts a literal (decimal, 0x hex, or 0b binary) into a uint64.
func parseNumber(literal string) (uint64, error) {
	lower := strings.ToLower(literal)
	switch {
	case strings.HasPrefix(lower, "0x"):
		return parseWithBase(literal, literal[2:], 16)
	case strings.HasPrefix(lower, "0b"):
		return parseWithBase(literal, literal[2:], 2)
	default:
		return parseWithBase(literal, literal, 10)
	}
}

func parseWithBase(literal, digits string, base int) (uint64, error) {
	value, err := strconv.ParseUint(digits, base, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number %q", literal)
	}
	return value, nil
}

// ---------------------------------------------------------------------------
// Parser (recursive descent, one method per precedence level)
// ---------------------------------------------------------------------------

type parser struct {
	tokens []token
	pos    int
}

func (p *parser) atEnd() bool    { return p.pos >= len(p.tokens) }
func (p *parser) peek() token    { return p.tokens[p.pos] }
func (p *parser) advance() token { t := p.tokens[p.pos]; p.pos++; return t }

// matchOp consumes and returns the next token if it is an operator matching one
// of the given texts.
func (p *parser) matchOp(ops ...string) (string, bool) {
	if p.atEnd() || p.peek().kind != tokOp {
		return "", false
	}
	for _, op := range ops {
		if p.peek().text == op {
			return p.advance().text, true
		}
	}
	return "", false
}

// Precedence, lowest to highest: | -> ^ -> & -> shift -> additive ->
// multiplicative -> unary -> primary.

func (p *parser) parseExpr() (uint64, error) { return p.parseOr() }

func (p *parser) parseOr() (uint64, error) {
	left, err := p.parseXor()
	if err != nil {
		return 0, err
	}
	for {
		if _, ok := p.matchOp("|"); !ok {
			return left, nil
		}
		right, err := p.parseXor()
		if err != nil {
			return 0, err
		}
		left |= right
	}
}

func (p *parser) parseXor() (uint64, error) {
	left, err := p.parseAnd()
	if err != nil {
		return 0, err
	}
	for {
		if _, ok := p.matchOp("^"); !ok {
			return left, nil
		}
		right, err := p.parseAnd()
		if err != nil {
			return 0, err
		}
		left ^= right
	}
}

func (p *parser) parseAnd() (uint64, error) {
	left, err := p.parseShift()
	if err != nil {
		return 0, err
	}
	for {
		if _, ok := p.matchOp("&"); !ok {
			return left, nil
		}
		right, err := p.parseShift()
		if err != nil {
			return 0, err
		}
		left &= right
	}
}

func (p *parser) parseShift() (uint64, error) {
	left, err := p.parseAdditive()
	if err != nil {
		return 0, err
	}
	for {
		op, ok := p.matchOp("<<", ">>")
		if !ok {
			return left, nil
		}
		right, err := p.parseAdditive()
		if err != nil {
			return 0, err
		}
		if op == "<<" {
			left <<= right
		} else {
			left >>= right
		}
	}
}

func (p *parser) parseAdditive() (uint64, error) {
	left, err := p.parseMultiplicative()
	if err != nil {
		return 0, err
	}
	for {
		op, ok := p.matchOp("+", "-")
		if !ok {
			return left, nil
		}
		right, err := p.parseMultiplicative()
		if err != nil {
			return 0, err
		}
		if op == "+" {
			left += right
		} else {
			left -= right
		}
	}
}

func (p *parser) parseMultiplicative() (uint64, error) {
	left, err := p.parseUnary()
	if err != nil {
		return 0, err
	}
	for {
		op, ok := p.matchOp("*", "/", "%")
		if !ok {
			return left, nil
		}
		right, err := p.parseUnary()
		if err != nil {
			return 0, err
		}
		switch op {
		case "*":
			left *= right
		case "/":
			if right == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			left /= right
		case "%":
			if right == 0 {
				return 0, fmt.Errorf("modulo by zero")
			}
			left %= right
		}
	}
}

func (p *parser) parseUnary() (uint64, error) {
	if op, ok := p.matchOp("~", "-", "+"); ok {
		operand, err := p.parseUnary()
		if err != nil {
			return 0, err
		}
		switch op {
		case "~":
			return ^operand, nil
		case "-":
			return -operand, nil // wraps in uint64, matching two's-complement
		default: // "+"
			return operand, nil
		}
	}
	return p.parsePrimary()
}

func (p *parser) parsePrimary() (uint64, error) {
	if p.atEnd() {
		return 0, fmt.Errorf("unexpected end of expression")
	}
	t := p.peek()
	switch t.kind {
	case tokNumber:
		p.advance()
		return t.value, nil
	case tokLParen:
		p.advance()
		value, err := p.parseExpr()
		if err != nil {
			return 0, err
		}
		if p.atEnd() || p.peek().kind != tokRParen {
			return 0, fmt.Errorf("missing closing parenthesis")
		}
		p.advance()
		return value, nil
	default:
		return 0, fmt.Errorf("unexpected token %q", t.text)
	}
}
