package piscine

func BasicAtoi(s string) int {
	num := 0
	for _, a:= range s {
		x := int(a - '0')
		num + num*10 + x
	}
	return num
}
