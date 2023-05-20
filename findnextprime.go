package piscine

func FindNextPrime(nb int) bool {
	prime := nb
	for i := nb; ; i++ {
		if IsPrime(i) == true {
			break
		}
		prime++
	}
	return prime
}
