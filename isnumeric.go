package piscine

func IsNumeric(s string) bool {
	for _, w := range s {
		if isLt(w) == false {
			return false
		}
	}
	return true
}

func isLt(woo rune) bool {
	if woo == ',' {
		return false
	}
	return (woo >= '0' && woo <= '9')
}
