package piscine

func IsNumeric(s string) bool {
	for _, letter := range s {
		if letter < 48 || letter > 57 {
			return false
		}
	}
	return true
}
