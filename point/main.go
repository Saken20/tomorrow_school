package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintInt(n int) {
	k := '0'
	for i := 0; i < n/10; i++ {
		k++
	}
	z01.PrintRune(k)

	nb1 := '0'
	for i := 0; i < n%10; i++ {
		nb1++
	}
	z01.PrintRune(nb1)
}

func main() {
	points := point{}
	setPoint(&points)

	s := "x = a, y = b\n"

	for _, ch := range s {
		if ch == 'a' {
			PrintInt(points.x)
		} else if ch == 'b' {
			PrintInt(points.y)
		} else {
			z01.PrintRune(ch)
		}
	}
}
