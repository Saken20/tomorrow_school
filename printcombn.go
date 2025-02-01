package piscine

import "github.com/01-edu/z01"

func generateComb(current string, n int, start int, isFirst *bool) {
	if len(current) == n {
		if !*isFirst {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		*isFirst = false
		for _, c := range current {
			z01.PrintRune(c)
		}
		return
	}
	for i := start; i < 10; i++ {
		generateComb(current+string('0'+byte(i)), n, i+1, isFirst)
	}
}

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	isFirst := true
	generateComb("", n, 0, &isFirst)
	z01.PrintRune('\n')
}
