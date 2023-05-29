package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	b := 1
	if n < 0 {
		b = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		c := (n / 10) * b
		if c != 0 {
			PrintNbr(c)
		}
		wee := (n % 10 * b) + '0'
		z01.PrintRune(rune(wee))
	} else {
		z01.PrintRune('0')
	}
}