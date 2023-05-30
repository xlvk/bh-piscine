package piscine

func SortCheck1(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := true
	descending := true

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			ascending = false
			break
		}
	}

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) < 0 {
			descending = false
			break
		}
	}

	return ascending || descending
}
