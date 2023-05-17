package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	sign := 1
	result := 0
	if s[0] == '+' || s[0] == '-' {
		if s[0] == '-' {
			sign = -1
		}
		s = s[1:]
	}
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0
		}
		digit := int(ch - '0')
		result = result*10 + digit
	}
	result *= sign
	return result
}
