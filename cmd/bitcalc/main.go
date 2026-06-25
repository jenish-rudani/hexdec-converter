// ============================================================================
//
// # Hexdec Converter
//
// File:        main.go
// Description: Main entry point for the bitwise/arithmetic calculator
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
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jenish-rudani/hexdec-converter/pkg/calculator"
)

var Version = "dev"

func main() {
	versionFlag := flag.Bool("version", false, "Print version information")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("hexdec-converter version %s\n", Version)
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Usage: bitcalc <expression>")
		fmt.Println("Example: bitcalc '1 << 2 & 0x03'")
		os.Exit(1)
	}

	// Allow either a single quoted expression or loose tokens; the shell still
	// requires quoting for metacharacters such as <<, &, |, and parentheses.
	expr := strings.Join(args, " ")

	result, err := calculator.Eval(expr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d\n", result)
	fmt.Printf("0x%X\n", result)
	fmt.Printf("0b%b\n", result)
}
