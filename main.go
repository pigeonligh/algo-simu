package main

import (
	"fmt"
	"os"

	"github.com/pigeonligh/algo-simu/simu"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: no algo kind\n")
		return
	}

	simu.Run(os.Args[1], os.Args[2:]...)
}
