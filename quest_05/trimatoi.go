package piscine

func TrimAtoi(s string) int {
	var num int
	var first_num bool
	for _, letter := range s {
		if letter >= '0' && letter <= '9' {
			num *= 10
			num += int(letter - 48)
		} else if letter == '-' && num == 0 {
			first_num = true
		}
	}
	if first_num {
		num = -num
	}
	return num
}
