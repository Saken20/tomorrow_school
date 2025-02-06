package piscine

func Map(f func(int) bool, a []int) []bool {
	tr := []bool{}

	for i, char := range a {
		tr[i] = f(int(char))
	}

	return tr
}
