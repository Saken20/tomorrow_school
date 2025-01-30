package piscine

func AppendRange(min, max int) []int {
	var n []int
	for i := min; i < max; i++ {
		n = append(n, i)
	}
	return n
}
