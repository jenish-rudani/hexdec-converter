// ============================================================================
//
// # Hexdec Converter
//
// File:        integration_test.go
// Description: Integration tests for the h2d and d2h command-line tools
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
package test

import (
	"os/exec"
	"strings"
	"testing"
)

func TestH2DIntegration(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"A", "10"},
		{"FF", "255"},
		{"100", "256"},
	}

	for _, test := range tests {
		cmd := exec.Command("go", "run", "../cmd/h2d/main.go", test.input)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Errorf("h2d failed for input %s: %v", test.input, err)
		}
		result := strings.TrimSpace(string(output))
		if result != test.expected {
			t.Errorf("h2d(%s) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestD2HIntegration(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"10", "A"},
		{"255", "FF"},
		{"256", "100"},
	}

	for _, test := range tests {
		cmd := exec.Command("go", "run", "../cmd/d2h/main.go", test.input)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Errorf("d2h failed for input %s: %v", test.input, err)
		}
		result := strings.TrimSpace(string(output))
		if result != test.expected {
			t.Errorf("d2h(%s) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestBitcalcIntegration(t *testing.T) {
	tests := []struct {
		input    string
		expected string // decimal, hex, binary lines joined by newline
	}{
		{"1 << 21", "2097152\n0x200000\n0b1000000000000000000000"},
		{"1 & 0x03", "1\n0x1\n0b1"},
		{"1 << 2 & 0x03", "0\n0x0\n0b0"},
	}

	for _, test := range tests {
		cmd := exec.Command("go", "run", "../cmd/bitcalc/main.go", test.input)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Errorf("bitcalc failed for input %q: %v", test.input, err)
		}
		result := strings.TrimSpace(string(output))
		if result != test.expected {
			t.Errorf("bitcalc(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}
