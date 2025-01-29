package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	var ans []rune
	if len(args) > 1 {
		ans = []rune(args[1])
	} else {
		ans = []rune(args[0])
	}
	for i := 2; i < len(ans); i++ {
		z01.PrintRune(rune(ans[i]))
	}
	z01.PrintRune('\n')
}
