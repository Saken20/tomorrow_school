package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for z := i; z <= '9'; z++ {
				for x := '0'; x <= '9'; x++ {
					if (i < z) || (i == z && j < x) {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(z)
						z01.PrintRune(x)
						if !(i == '9' && j == '8' && z == '9' && x == '9') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}