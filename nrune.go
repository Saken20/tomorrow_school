package piscine

func NRune(s string, n int) rune {
	R := []rune(s)
	if n < 1 || len(R) < n {
		return 0
	}
	return R[n-1]
}
