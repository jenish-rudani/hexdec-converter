// ============================================================================
//
// # Hexdec Converter
//
// File:        main.go
// Description: Main entry point for the hexadecimal to decimal converter
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
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jenish-rudani/hexdec-converter/pkg/converter"
)

var Version = "dev"

func main() {
	versionFlag := flag.Bool("version", false, "Print version information")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("hexdec-converter version %s\n", Version)
		return
	}
	if len(os.Args) != 2 {
		fmt.Println("Usage: h2d <hexadecimal number>")
		os.Exit(1)
	}

	hex := os.Args[1]
	decimal, err := converter.HexToDecimal(hex)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d\n", decimal)
}
