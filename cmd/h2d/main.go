package main

import (
    "fmt"
    "os"

    "github.com/jenish-rudani/hexdec-converter/pkg/converter"
)

func main() {
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
