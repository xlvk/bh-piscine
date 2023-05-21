package piscine

func IsUpper(s string) bool {
	for _, w := range s {
		if isLette(w) == false {
			return false
		}
	}
	return true
}

func isLette(woo rune) bool {
	return (woo >= 'A' && woo <= 'Z')
}
