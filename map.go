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

ffunc Map(f func(int) bool, a []int) []bool {
	m := make(map[int]bool)
	for _, i := range a {
		m[i] = f(i) 
	}
	wee := make([]bool, len(m))
	for o, j := range wee {
		wee[o] = j
	}
	return wee
}
