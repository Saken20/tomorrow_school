package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x, y int
}

func prns(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func prni(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		prni(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := point{}
	setPoint(&points)

	prns("x = ")
	prni(points.x)
	prns(", y = ")
	prni(points.y)
	prns('\n')
}
