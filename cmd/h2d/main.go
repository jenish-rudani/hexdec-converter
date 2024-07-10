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
