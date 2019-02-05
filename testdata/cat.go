// Package main provides a cat like utility that echos standard in to standard out.
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
