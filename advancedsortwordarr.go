package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	n := len(a)
	if n <= 1 {
		return
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
