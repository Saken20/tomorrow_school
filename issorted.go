package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	inc := true // өсу ретімен
	dec := true // кему ретімен
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			inc = false
		} else if f(a[i], a[i+1]) > 0 {
			dec = false
		}
	}
	return !(!inc && !dec)
}
