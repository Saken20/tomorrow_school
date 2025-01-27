package piscine

func RecursiveFactorial(nb int) int {
	if nb < 1 {
		return 0
	} else {
		return nb * RecursiveFactorial(nb-1)
	}
}
