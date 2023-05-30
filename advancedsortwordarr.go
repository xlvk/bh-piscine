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
	i := 0
	for i < len(a) && i < len(b) {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
		i++
	}
	// If the common prefix is the same, the shorter string should come first
	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}
	return 0
}
