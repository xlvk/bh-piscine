package piscine

func AlphaCount(s string) int {
	count := 0
	for _, char := range s {
		if isLetter(char) {
			count++
		}
	}
	return count
}

func isLetter(char rune) bool {
	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z')
}
