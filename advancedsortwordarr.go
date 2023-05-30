package piscine

import "strings"

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			// Swap elements if the comparison result is greater than 0
			a[i-1], a[i] = a[i], a[i-1]
		}
	}
}

func wawawawa(a, b string) int {
	return strings.Compare(a, b)
}
