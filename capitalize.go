package piscine

func Capitalize(s string) string {
	var t string = ""
	if len(s) == 0 {
		return t
	} else if IsLet(s[0]) == 1 {
		t += string(s[0] - 32)
	} else {
		t += string(s[0])
	}
	for i := 1; i < len(s); i++ {
		if IsLet(s[i]) == 2 {
			if s[i-1] != ' ' && IsLet(s[i-1]) != 0 {
				t += string(s[i] + 32)
			} else {
				t += string(s[i])
			}
		} else if IsLet(s[i]) == 1 {
			if s[i-1] == ' ' || IsLet(s[i-1]) == 0 {
				t += string(s[i] - 32)
			} else {
				t += string(s[i])
			}
		} else {
			t += string(s[i])
		}
	}
	return t
}

func IsLet(s byte) int {
	if s >= '0' && s <= '9' {
		return 3
	} else if s >= 'A' && s <= 'Z' {
		return 2
	} else if s >= 'a' && s <= 'z' {
		return 1
	}
	return 0
}
