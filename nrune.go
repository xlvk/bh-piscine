package piscine

func NRune(s string, n int) rune {
	if n <= 0 || len(s) == 0 || n > len([]rune(s)) {
		return 0
	}

	return []rune(s)[n-1]
}
