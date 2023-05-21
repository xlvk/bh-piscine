package piscine

func IsLower(s string) bool {
	for _, w := range s {
		if isLete(w) == false {
			return false
		}
	}
	return true
}

func isLete(woo rune) bool {
	return (woo >= 'a' && woo <= 'z')
}
