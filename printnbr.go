package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int64) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		PrintNbr(n / 10)
	}
	digit := '0' + n%10
	z01.PrintRune(digit)
}
