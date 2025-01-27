package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = n * -1
	}
	for {
		digit := n % 10
		if digit < 0 {
			digit = digit * -1
		}
		defer z01.PrintRune(rune(digit) + '0')
		n = n / 10
		if n == 0 {
			break
		}
	}
}
