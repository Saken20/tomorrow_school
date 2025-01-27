package piscine

func BasicAtoi2(s string) int {
	var x int = 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			x = x*10 + int(char-48)
		} else {
			return 0
		}
	}
	return x
}
