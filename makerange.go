package piscine

func MakeRange(min, max int) []int {
	var n []int
	if min >= max {
		return n
	}
	size_n := max - min
	n = make([]int, size_n)
	for i := 0; i < size_n; i++ {
		n[i] = i + min
	}
	return n
}
