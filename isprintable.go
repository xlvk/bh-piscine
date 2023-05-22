package piscine

func IsPrintable(s string) bool {
	for _, r := range s {
		if !isPrintabe(r) {
			return false
		}
	}
	return true
}

func isPrintabe(r rune) bool {
	return !(r == '\n' || (r >= '0' && r <= '9'))
}
