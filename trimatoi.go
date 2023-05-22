package piscine

func TrimAtoi(w string) int {
	size := 0
	ind := '0'
	var s string
	for p := range w {
		if rune(w[p]) >= '0' && rune(w[p]) <= '9' {
			break
		}
		if rune(w[p]) == '+' {
			ind = '+'
			break
		}
		if rune(w[p]) == '-' {
			ind = '-'
			break
		}
	}
	for p := range w {
		if rune(w[p]) >= '0' && rune(w[p]) <= '9' {
			s += string(w[p])
			size++
		}
	}
	d := 1
	num := 0
	for i := size - 1; i >= 0; i-- {
		num += d * int(s[i]-'0')
		d *= 10
	}
	if ind == '0' || ind == '+' {
		return num
	} else {
		return -num
	}
}
