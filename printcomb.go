package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	isstartt := true
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for k := j + 1; k <= '9'; k++ {
				if !isstartt {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				} else {
					isstartt = false
				}
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
			}
		}
	}
	z01.PrintRune('\n')
}
