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
