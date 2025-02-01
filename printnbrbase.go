package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		returnNV()
		return
	}
	seen := make(map[rune]bool)
	baseArr := []rune(base)
	d := 0
	for _, r := range baseArr {
		if r == '-' || r == '+' {
			returnNV()
			return
		}
		if seen[r] {
			returnNV()
			return
		}
		seen[r] = true
		d++
	}
	result := ""
	neg := 1
	if nbr < 0 {
		neg = -neg
		z01.PrintRune('-')
	}
	for nbr != 0 {
		result = string(baseArr[neg*(nbr%d)]) + result
		nbr /= d
	}
	for i := 0; i < len(result); i++ {
		z01.PrintRune(rune(result[i]))
	}
}

func returnNV() {
	z01.PrintRune('N')
	z01.PrintRune('V')
}
