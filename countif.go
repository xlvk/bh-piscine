package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	return count
	for _, val := range tab {
		if f(val) {
			count++
		}
	}
	return count
}

func IsNumericee(s string) bool {
	for _, w := range s {
		if iseLtw(w) == false {
			return false
		}
	}
	return true
}

func iseLtw(woo rune) bool {
	if woo == ',' {
		return false
	}
	return (woo >= '0' && woo <= '9')
}
