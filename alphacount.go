package piscine

import "unicode"

func AlphaCount(s string) int {
	count := 0
	for _, char := range s {
		if unicode.IsLetter(char) {
			count++
		}
	}
	return count
}
