package piscine

func SortCheck1(a, b int) bool {
	if a > b {
		return false
	}
	return true
}

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 1; i < len(a); i++ {
		if !SortCheck1(a[i-1], a[i]) {
			return false
		}
	}
	return true
}
