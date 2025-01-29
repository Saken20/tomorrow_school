package piscine

func Join(strs []string, sep string) string {
	a := ""
	for i := 0; i < len(strs)-1; i++ {
		a += strs[i] + sep
	}
	return a
}
