package piscine

func IsPrintable(s string) bool {
	for _, letter := range s {
		if letter < 65 || (letter > 90 && letter < 97) || letter > 122 {
			return false
		}
	}
	return true
}
