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
        {"FFFFFFFFFFFFFFFFFFF", 0, true},  // Too large
        {"G", 0, true},  // Invalid hex character
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