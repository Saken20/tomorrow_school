package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, val := range a {
		PrintStr(val)
	}
}

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
