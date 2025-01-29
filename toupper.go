package piscine

func ToUpper(s string) string {
	str := ""
	for _, letter := range s {
		if letter >= 97 || letter <= 122 {
			str += string(letter - 32)
		} else {
			str += string(letter)
		}
	}
	return str
}
