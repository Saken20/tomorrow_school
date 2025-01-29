package piscine

func IsUpper(s string) bool {
	b := true
	for _, letter := range s {
		if letter < 65 && letter > 90 {
			b = false
		}
	}
	return b
}
