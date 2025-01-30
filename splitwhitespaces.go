package piscine

func SplitWhiteSpaces(s string) []string {
	var ans []string
	var word string
	for _, ch := range s {
		if ch != ' ' {
			word += string(ch)
		} else {
			if word != "" {
				ans = append(ans, word)
			}
			word = ""
		}
	}
	if word != "" {
		ans = append(ans, word)
	}
	return ans
}
