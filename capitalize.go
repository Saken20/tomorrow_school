package piscine

func Capitalize(s string) string {
	a := ToUpper(string(s[0]))
	if len(s) == 0 {
		return a
	} else {
		for i := 1; i < len(s); i++ {
			if (s[i-1] >= '0' && s[i-1] <= '9') || (s[i-1] >= 'A' && s[i-1] <= 'Z') || (s[i-1] >= 'a' && s[i-1] <= 'z') {
				a += ToLower(string(s[i]))
			} else if (s[i-1] < '0' && s[i-1] > '9') || (s[i-1] < 'A' && s[i-1] > 'Z') || (s[i-1] < 'a' && s[i-1] > 'z') {
				a += ToUpper(string(s[i]))
			} else {
				a += string(s[i])
			}
		}
	}
	return a
}
