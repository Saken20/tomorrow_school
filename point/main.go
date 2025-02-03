package main

import (
	"fmt"
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

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
