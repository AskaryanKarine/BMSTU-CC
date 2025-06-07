package main

import (
	"flag"
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/compiler"
	"os"
)

var (
	input  = flag.String("i", "examples/2+2.kum", "source .kum file")
	output = flag.String("o", "examples/out", "executable program")
	isDOT  = flag.Bool("d", false, "output as DOT")
)

func main() {
	flag.Parse()

	err := compiler.Compiler(*input, *output, *isDOT)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
