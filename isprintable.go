package piscine

func IsPrintable(s string) bool {
	for _, w := range s {
		if isL(w) == false {
			return false
		}
	}
	return true
}

func isL(woo rune) bool {
	if woo == '\n' {
		return false
	}
	return (woo >= '0' && woo <= '9')
}
