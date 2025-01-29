package piscine

func BasicJoin(elems []string) string {
	a := ""
	for i := 0; i < len(elems); i++ {
		a += elems[i]
	}
	return a
}
