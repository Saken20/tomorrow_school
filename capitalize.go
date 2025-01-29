package piscine

func Capitalize(s string) string {
	a := ToUpper(string(s[0]))
	if len(s) == 0 {
		return a
	} else {
		for i := 1; i < len(s); i++ {
			if s[i-1] >= 32 && s[i-1] <= 47 {
				a += ToUpper(string(s[i]))
			} else if (s[i-1] >= 65 && s[i-1] <= 90) || (s[i-1] >= 48 && s[i-1] <= 57) {
				a += ToLower(string(s[i]))
			} else {
				a += string(s[i])
			}
		}
	}
	return a
}
