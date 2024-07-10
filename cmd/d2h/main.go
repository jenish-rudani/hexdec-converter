package main

import (
    "fmt"
    "os"
    "strconv"

    "github.com/jenish-rudani/hexdec-converter/pkg/converter"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: d2h <decimal number>")
        os.Exit(1)
    }

    decimal, err := strconv.ParseUint(os.Args[1], 10, 64)
    if err != nil {
        fmt.Printf("Error: Invalid decimal number '%s'\n", os.Args[1])
        os.Exit(1)
    }

    hex := converter.DecimalToHex(decimal)
    fmt.Printf("%s\n", hex)
}
