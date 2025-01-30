package piscine

func Split(s, sep string) []string {
	var ans []string
	var word string
	sepLen := len(sep)
	for i := 0; i < len(s); i++ {
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			if word != "" {
				ans = append(ans, word)
				word = ""
			}
			i += sepLen - 1
		} else {
			word += string(s[i])
		}
	}
	if word != "" {
		ans = append(ans, word)
	}
	return ans
}
