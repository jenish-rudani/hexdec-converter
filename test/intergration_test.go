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