package piscine

func IsPrime(num int) bool {
	if num == 1 || num == 0 {
		return false
	}

	// Check divisibility up to the square root of the number
	limit := intSqrte(num)
	for i := 2; i <= limit; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func intSqrte(num int) int {
	// Use Newton's method to find the integer square root
	x := num
	y := (x + 1) / 2
	for y < x {
		x = y
		y = (x + num/x) / 2
	}
	return x
}

func Map(f func(int) bool, a []int) []bool {
	m := make(map[int]bool)
	for _, i := range a {
		m[i] = f(i)
	}
	wee := make([]bool, len(m))
	for o, j := range wee {
		wee[o] = j
	}
	return wee
}
