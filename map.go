package piscine

func IsPrime(nb int) bool {
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

func Map(f func(int) bool, a []int) []bool {
	for _, i := range a {
		f(i)
	}
	return
}
