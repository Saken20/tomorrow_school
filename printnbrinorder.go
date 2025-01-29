package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	var t []int
	if n == 0 {
		t = append(t, 0)
	}
	for n > 0 {
		t = append(t, n%10)
		n /= 10
	}
	SortIntegerTable(t)
	for _, val := range t {
		z01.PrintRune(rune(val) + '0')
	}
}

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := i + 1; j < len(table); j++ {
			if table[j] < table[i] {
				table[j], table[i] = table[i], table[j]
			}
		}
	}
}
