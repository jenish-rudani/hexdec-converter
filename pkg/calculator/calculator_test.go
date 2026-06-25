// ============================================================================
//
// # Hexdec Converter
//
// File:        calculator_test.go
// Description: Unit tests for the bitwise/arithmetic expression calculator
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
package calculator

import (
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr     string
		expected uint64
	}{
		// Literals
		{"42", 42},
		{"0x03", 3},
		{"0X1F", 31},
		{"0b1010", 10},
		{"0B1010", 10},
		{"0", 0},

		// Single operators
		{"1 << 21", 2097152},
		{"256 >> 4", 16},
		{"1 & 0x03", 1},
		{"0x0F | 0xF0", 255},
		{"0xFF ^ 0x0F", 240},
		{"2 + 3", 5},
		{"10 - 4", 6},
		{"6 * 7", 42},
		{"20 / 4", 5},
		{"17 % 5", 2},

		// Unary
		{"~0", 18446744073709551615},
		{"~0xFF", 18446744073709551360},
		{"-1", 18446744073709551615},
		{"+5", 5},

		// Precedence (C-style)
		{"1 << 2 & 0x03", 0}, // (1<<2) & 3 = 4 & 3 = 0
		{"1 & 0x03 << 2", 0}, // 1 & (3<<2) = 1 & 12 = 0
		{"2 + 3 * 4", 14},    // 2 + 12
		{"1 | 2 & 3", 3},     // 1 | (2 & 3) = 1 | 2 = 3
		{"1 ^ 1 << 1", 3},    // 1 ^ (1<<1) = 1 ^ 2 = 3
		{"100 - 10 - 5", 85}, // left-associative

		// Parentheses
		{"(1 + 2) * 3", 9},
		{"1 & (0x03 << 2)", 0},
		{"(2 + 3) * 4", 20},

		// Whitespace insensitivity
		{"1<<21", 2097152},
		{"  1  +  2  ", 3},

		// Wraparound
		{"0 - 1", 18446744073709551615},
	}

	for _, test := range tests {
		result, err := Eval(test.expr)
		if err != nil {
			t.Errorf("Eval(%q) returned unexpected error: %v", test.expr, err)
			continue
		}
		if result != test.expected {
			t.Errorf("Eval(%q) = %d, expected %d", test.expr, result, test.expected)
		}
	}
}

func TestEvalErrors(t *testing.T) {
	tests := []string{
		"",       // empty
		"   ",    // whitespace only
		"1 +",    // trailing operator
		"* 2",    // leading binary operator
		"(1 + 2", // unbalanced open paren
		"1 + 2)", // unbalanced close paren
		"1 / 0",  // division by zero
		"1 % 0",  // modulo by zero
		"1 2",    // two numbers, no operator
		"0xZZ",   // invalid hex
		"0b12",   // invalid binary
		"@",      // unknown character
	}

	for _, expr := range tests {
		_, err := Eval(expr)
		if err == nil {
			t.Errorf("Eval(%q) expected an error, but got none", expr)
		}
	}
}
