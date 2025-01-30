package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Convert(s string) int {
	n := 0
	for _, val := range s {
		if val >= 48 && val <= 57 {
			n = n*10 + int(val-48)
		}
	}
	if n < 1 || n > 26 {
		return -64
	}
	return n
}

func main() {
	var s []string
	j := 0
	if len(os.Args) == 1 {
		return
	} else if os.Args[1] == "--upper" {
		j = 32
		s = append(s, os.Args[2:]...)
	} else {
		s = append(s, os.Args[1:]...)
	}
	for _, val := range s {
		z01.PrintRune(rune(Convert(val) + 96 - j))
	}
	z01.PrintRune('\n')
}
