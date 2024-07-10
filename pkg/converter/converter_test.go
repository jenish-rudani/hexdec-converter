// ============================================================================
//
// # Hexdec Converter
//
// File:        converter_test.go
// Description: Unit tests for the converter package
//
// Author:      Jenish Rudani
// Created:     2024-07-10
// Updated:     2024-07-10
//
// Version:     1.0
// Copyright:   (c) 2024 Jenish Rudani
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
package converter

import (
	"testing"
)

func TestHexToDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected uint64
		hasError bool
	}{
		{"A", 10, false},
		{"F", 15, false},
		{"10", 16, false},
		{"FF", 255, false},
		{"100", 256, false},
		{"FFFFFFFFFFFFFFFF", 18446744073709551615, false},
		{"FFFFFFFFFFFFFFFFFFF", 0, true}, // Too large
		{"G", 0, true},                   // Invalid hex character
	}

	for _, test := range tests {
		result, err := HexToDecimal(test.input)
		if test.hasError {
			if err == nil {
				t.Errorf("HexToDecimal(%s) expected an error, but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("HexToDecimal(%s) returned unexpected error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("HexToDecimal(%s) = %d, expected %d", test.input, result, test.expected)
			}
		}
	}
}

func TestDecimalToHex(t *testing.T) {
	tests := []struct {
		input    uint64
		expected string
	}{
		{10, "A"},
		{15, "F"},
		{16, "10"},
		{255, "FF"},
		{256, "100"},
		{18446744073709551615, "FFFFFFFFFFFFFFFF"},
	}

	for _, test := range tests {
		result := DecimalToHex(test.input)
		if result != test.expected {
			t.Errorf("DecimalToHex(%d) = %s, expected %s", test.input, result, test.expected)
		}
	}
}
