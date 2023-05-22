package piscine

func TrimAtoi(s string) int {
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
			b := 0
			for i := 0; i < int(ch-'0'); i++ {
				b++
			}
			result = result*10 + b
		}
	}
	result *= sign
	return result
}
