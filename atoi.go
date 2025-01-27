package piscine

func Atoi(s string) int {
	newStr := []rune(s)
	count := StrLen(s)
	sstr := ""
	if s == "+" || s == "-" || s == "" {
		return 0
	}
	for i := 1; i < count; i++ {
		sstr = sstr + string(newStr[i])
	}
	if newStr[0] == '-' {
		return -1 * BasicAtoi2(sstr)
	}
	if newStr[0] == '+' {
		return BasicAtoi2(sstr)
	}
	return BasicAtoi2(s)
}

func BasicAtoi2(s string) int {
	count := StrLen(s)
	newStr := []rune(s)
	val := 0
	for i := 0; i < count; i++ {
		t := 1
		for j := 0; j < count-i-1; j++ {
			t = t * 10
		}
		sym := 0
		for k := '0'; k < newStr[i]; k++ {
			sym++
		}
		if sym == 0 && newStr[i] != '0' {
			return 0
		}
		if sym > 9 {
			return 0
		}

		val = val + sym*t
	}
	return val
}

func StrLen(str string) int {
	newStr := []rune(str)
	count := 0
	for range newStr {
		count++
	}
	return count
}
