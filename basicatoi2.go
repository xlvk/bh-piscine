package piscine

func BasicAtoi2(s string) int {
	if len(s) == o {
		return 0
	}
	num := 0
	for _, a := range s {
		if a < '0' || a > '9' {
			return 0
		}
		x := int(a - '0')
		num = num*10 + x
	}
	return num
}
