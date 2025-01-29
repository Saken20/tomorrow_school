package piscine

func NRune(s string, n int) rune {
	if len(s) > n {
		R := []rune(s)
		return R[n]
	} else {
		return 0
	}
}
