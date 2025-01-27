package piscine

func RecursiveFactorial(nb int) int {
	if nb < 1 {
		return 0
	}
	if nb > 0  && nb < 1000{
		return nb * RecursiveFactorial(nb-1)
	}
	return 0
}
