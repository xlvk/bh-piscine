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
loop:
	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if f(a[j], a[i]) > 0 {
				continue loop
			}
		}
		return false
	}
	return true
}
