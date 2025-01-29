package piscine

func Join(strs []string, sep string) string {
	a := ""
	for i := 0; i < len(strs); i++ {
		if i < len(strs) - 1 {
			a += strs[i] + sep
		} else {
			a += strs[i]
		}
	}
	return a
}
