package main

import "github.com/01-edu/z01"

type point struct {
	x, y int
}

func main() {
	points := &point{}
	setPoint(points)

	printStr("x = ")
	printNumber(points.x)
	printStr(", y = ")
	printNumber(points.y)
	z01.PrintRune('\n')
}

func printStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func printNumber(nbr int) {
	digits := [20]int{}
	i := 0

	for nbr > 0 {
		digit := nbr % 10
		digits[i] = digit
		nbr /= 10
		i++
	}

	for j := i - 1; j >= 0; j-- {
		index := '0'
		for k := 0; k < digits[j]; k++ {
			index++
		}
		z01.PrintRune(index)
	}
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}
