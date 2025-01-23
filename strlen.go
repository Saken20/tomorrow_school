package piscine

func StrLen(s string) int {
	var counter int = 0
	for _, char := range s {
		if char <= 255 {
			counter ++
		}
	}
	return counter
}
