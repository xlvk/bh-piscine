package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune(0)
		return
	}
	if n > 0 {
	}
	digits := []int{}
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}
	reverseSlice(digits)
	return digits
}

func reverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
