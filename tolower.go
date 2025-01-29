package piscine

func ToLower(s string) string {
	str := ""
	for i, letter := range s {
		if letter >= 'A' && letter <= 'Z' {
			str += string(s[i] + 32)
		} else {
			str += string(s[i])
		}
	}
	return str
}