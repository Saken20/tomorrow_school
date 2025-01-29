package piscine

func ToUpper(s string) string {
	str := ""
	for i, letter := range s {
		if letter >= 'a' || letter <= 'z' {
			str += string(s[i] - 32)
		} else {
			str += string(s[i])
		}
	}
	return str
}
