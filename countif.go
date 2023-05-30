package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, str := range tab {
		if f(str) {
			count++
		}
	}
	return count
}

func IsNumericw(s string) bool {
	for _, char := range s {
		if !isDigit(char) {
			return false
		}
	}
	return true
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
