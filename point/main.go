package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	SolveNbr(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	SolveNbr(points.y)
	z01.PrintRune('\n')
}

func SolveNbr(n int) {
	b := 1
	if n < 0 {
		b = -1
		z01.PrintRune('-')
		n = -n
	}
	if n != 0 {
		digits := make([]int, 0)
		for n > 0 {
			digit := n % 10
			digits = append(digits, digit)
			n /= 10
		}
		for i := len(digits) - 1; i >= 0; i-- {
			z01.PrintRune('0' + digits[i]*b)
		}
	} else {
		z01.PrintRune('0')
	}
}
