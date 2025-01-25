package main

import (
	"fmt"
)

func printRune(r rune) {
	fmt.Print(string(r))
}

const (
	lu = 0
	ru = 1
	ld = 2
	rd = 3
	u  = 4
	l  = 5
	r  = 6
	d  = 7
)

//	0	 1     2    3    4   5   6   7
//  'lu' 'ru' 'ld' 'rd' 'u' 'l' 'r' 'd'

func mainQuad(x int, y int, sym []rune) {
	if x < 1 && y < 1 {
		return
	}

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if col < 1 && row < 1 {
				printRune(sym[lu]) // left up
			} else if col == x-1 && row < 1 {
				printRune(sym[ru]) // right up
			} else if col < 1 && row == y-1 {
				printRune(sym[ld]) // left down
			} else if col == x-1 && row == y-1 {
				printRune(sym[rd]) // right down
			} else if col < 1 {
				printRune(sym[l]) // left
			} else if row < 1 {
				printRune(sym[u]) // up
			} else if col == x-1 {
				printRune(sym[r]) // right
			} else if row == y-1 {
				printRune(sym[d]) // down
			} else {
				printRune(' ')
			}
		}
		printRune('\n')
	}
}

//	0	 1     2    3    4   5   6   7
//  'lu' 'ru' 'ld' 'rd' 'u' 'l' 'r' 'd'

func main() {
	mainQuad(10, 10, []rune{'/', '\\', '\\', '/', '-', '|', '|', '_'})
}
