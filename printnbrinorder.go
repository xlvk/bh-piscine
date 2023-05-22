package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune(0)
		return
	}

	digits := []int{}
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}

	reverseSlice(digits)

	for _, digi := range bubbleSort(digits) {
		z01.PrintRune(digi)
	}
}

func reverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func bubbleSort(slice []int) []int {
	length := len(slice)
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < length-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				swapped = true
			}
		}
	}

	return slice
}
