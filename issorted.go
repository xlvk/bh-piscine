package piscine

func SortIntegerTablee(table []int) {
	n := len(table)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			return false
		}
	}
	return true
}

func isSortedBy10(a, b int) int {
	return (a / 10) - (b / 10)
}
