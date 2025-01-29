package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		return
	}

	args = args[len(args)-1:1]

	for _, word := range args {
		for _, ch := range word {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}