package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"github.com/pigeonligh/algo-simu/algo/factory"
)

type fdWriter struct {
	fd int
}

func (c *fdWriter) Write(p []byte) (n int, err error) {
	return syscall.Write(c.fd, p)
}

func init() {
	factory.Register()
}

func main() {
	output := &fdWriter{3}
	scanner := bufio.NewScanner(os.Stdin)

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: no algo kind\n")
		fmt.Fprintf(output, "end error\n")
		return
	}

	algo, err := factory.New(os.Args[1], func(value string) string {
		fmt.Fprintf(output, "check %v\n", value)
		if scanner.Scan() {
			return scanner.Text()
		}

		fmt.Fprintf(output, "end error\n")
		panic("input closed")
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(output, "end error\n")
		return
	}

	if err := algo.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(output, "end error\n")
		return
	}

	if err := algo.Solve(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(output, "end error\n")
		return
	}

	fmt.Fprintf(output, "end %v\n", algo.Result())
}
