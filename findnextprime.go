package piscine

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb >= 100 {
		return -1
	}
	if IsPrime(nb) {
		return nb
	}
	return FindNextPrime(nb + 1)
}