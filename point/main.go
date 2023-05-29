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

func PrintStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func main() {
	points := &point{}
	setPoint(points)
	PrintStr("x = ")
	n := points.x
	if n < 0 {
		PrintStr("-")
		n = -n
	}
	printDigits(n)
	PrintStr(", y = ")
	b := 1
	n = points.y
	if n < 0 {
		b = -1
		PrintStr("-")
		n = -n
	}
	printDigits(n)
	z01.PrintRune('\n')
}

func printDigits(n int) {
	digits := [20]int{}
	i := 0
	for n > 0 {
		digit := n % 10
		digits[i] = digit
		n /= 10
		i++
	}
	power := 1
	for j := i - 1; j >= 0; j-- {
		digit := digits[j]
		for k := 0; k < power; k++ {
			digit *= 10
		}
		z01.PrintRune('0' + digit)
		power *= 10
	}
}
