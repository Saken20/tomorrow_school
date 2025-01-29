package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var s []string
	s = append(s, os.Args[1:]...)
	SortIntegerTable(s)
	for _, val := range s {
		Print([]rune(val))
	}
}

func Print(ans []rune) {
	for i := 0; i < len(ans); i++ {
		z01.PrintRune(rune(ans[i]))
	}
	z01.PrintRune('\n')
}

func SortIntegerTable(table []string) {
	for i := 0; i < len(table); i++ {
		for j := i + 1; j < len(table); j++ {
			if []rune(table[j])[0] < []rune(table[i])[0] {
				table[j], table[i] = table[i], table[j]
			}
		}
	}
}
