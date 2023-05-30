package piscine

func Any(f func(string) bool, a []string) bool {
	for i, val := range a {
		if f(val) {
			return true
		}
		if i == len(a)-1 {
			return false
		}
	}
	return false
}

func IsNumerice(s string) bool {
	for _, w := range s {
		if iseLt(w) == false {
			return false
		}
	}
	return true
}

func iseLt(woo rune) bool {
	if woo == ',' {
		return false
	}
	return (woo >= '0' && woo <= '9')
}
