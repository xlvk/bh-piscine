package piscine

func IsItPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) bool {
	if nb < 2 {
		return 2
	}
	for i := nb + 1; ; i++ {
		if IsItPrime(i) {
			prime := i
		}
	}
	return prime
}
