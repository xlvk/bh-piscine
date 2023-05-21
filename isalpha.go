package piscine

func IsAlpha(s string) bool {
	for _, w := range s {
		if w == ' ' {
			return false
		}
	}
	return true
}
