package piscine

func Map(f func(int) bool, a []int) []bool {
	tr := []bool{}

	for _, char := range a {
		tr = append(tr, f(char))
	}

	return tr
}
