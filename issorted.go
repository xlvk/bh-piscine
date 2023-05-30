package piscine

func SortCheck1(a, b int) bool {
	return a <= b
}

func IsSorted(f func(a, b int) bool, a []int) bool {
	ascending := true
	descending := true
	for i := 1; i < len(a); i++ {
		if !f(a[i-1], a[i]) {
			ascending = false
			break
		}
	}

	for i := 1; i < len(a); i++ {
		if !f(a[i], a[i-1]) {
			descending = false
			break
		}
	}

	return ascending || descending
}
