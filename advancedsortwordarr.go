package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			// Swap elements if the comparison result is greater than 0
			a[i-1], a[i] = a[i], a[i-1]
		}
	}
}

func wawawawa(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
