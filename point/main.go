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

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := point{}
	setPoint(&points)

	printStr("x = ")
	for _, r := range string(points.x) {
		z01.PrintRune(r)
	}
	printStr(", y = ")
	for _, k := range string(points.y) {
		z01.PrintRune(k)
	}
	z01.PrintRune('\n')
}
