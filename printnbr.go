package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int64) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	digits := make([]int64, 0)
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(rune(digits[i]) + '0')
	}
}
