package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	digits := [9]int{}
	for i := 0; i < n; i++ {
		digits[i] = i
	}
	for {
		printDigits(digits[:n])
		// Find the rightmost digit to increment
		i := n - 1
		for i >= 0 && digits[i] == 10-n+i {
			i--
		}
		if i < 0 {
			break // All combinations printed
		}
		digits[i]++
		for j := i + 1; j < n; j++ {
			digits[j] = digits[j-1] + 1
		}
	}
}

// Helper function to print the digits combination
func printDigits(digits []int) {
	for _, digit := range digits {
		z01.PrintRune(rune(digit) + '0')
	}
	if digits[0] != 10-len(digits) {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	} else {
		z01.PrintRune('\n')
	}
}