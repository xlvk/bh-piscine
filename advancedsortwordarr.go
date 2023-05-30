package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if f(a[i], a[j]) > 0 {
				// Swap elements if the comparison result is greater than 0
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func wowowowwo(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
