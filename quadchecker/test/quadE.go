package main

import (
	"fmt"
	"piscine"
)

func main() {
	x, y, err := piscine.GetNumber()
	if err != nil {
		return
	}
	fmt.Print(piscine.QuadE(x, y))
}
