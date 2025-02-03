package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x, y int
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printInt(n int) {
	k := '0'
	for i := 0; i < n / 10; i++ {
		k++
	}
	z01.PrintRune(k)
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := point{}
	setPoint(&points)

	s := "x = a, y = b\n"

	for _, ch := range s {
		if ch == 'a' {
			printInt(point.x)
		} else if ch == 'b' {
			z01.PrintRune('b')
		} else {
			z01.PrintRune(ch)
		}
	}
}
