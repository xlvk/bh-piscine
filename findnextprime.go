package piscine

func FindNextPrime(nb int) int {
	poo := nb
	for i := nb; ; i++ {
		if IsPrime(i) == true {
			break
		}
		poo++
	}
	return poo
}
