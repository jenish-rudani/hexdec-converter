// ============================================================================
//
// # Hexdec Converter
//
// File:        converter.go
// Description: Core logic for hexadecimal and decimal conversions
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
	"fmt"
	"strconv"
)

// HexToDecimal converts a hexadecimal string to a uint64 decimal
func HexToDecimal(hex string) (uint64, error) {
	return strconv.ParseUint(hex, 16, 64)
}

// DecimalToHex converts a uint64 decimal to a hexadecimal string
func DecimalToHex(decimal uint64) string {
	return fmt.Sprintf("%X", decimal)
}
