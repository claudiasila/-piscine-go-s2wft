package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}

	for {
		if isPrime(nb) {
			return nb
		}
		nb++
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}