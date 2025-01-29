package piscine

func LastRune(s string) rune {
	R := []rune(s)
	Rlen := len(s) - 1
	return R[Rlen]
}
