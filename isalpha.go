package piscine

func IsAlpha(s string) bool {
	for _, letter := range s {
		if letter < 48 || (letter > 57 && letter < 65) || (letter > 90 && letter < 97) || letter > 122 {
			return false
		}
	}
	return true
}
