package piscine

func AtoiBase(s string, base string) int {
	if len(base) < 2 {
		return 0
	}
	boy := make(map[rune]bool)
	baseArr := []rune(base)
	d := 0
	hash := make(map[rune]int)
	for i, r := range baseArr {
		if r == '-' || r == '+' {
			return 0
		}
		if boy[r] {
			return 0
		}
		hash[r] = i
		boy[r] = true
		d++
	}
	result := 0
	sp := 0
	for i := len(s) - 1; i > -1; i-- {
		result += hash[rune(s[i])] * pow(d, sp)
		sp++
	}
	return result
}

func pow(n, g int) int {
	result := 1
	for i := 0; i < g; i++ {
		result *= n
	}
	return result
}
