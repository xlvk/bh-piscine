package piscine

func SortCheck1(a, b int) bool {
	if a > b {
		return false
	}
	return true
}

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 1; i < len(a); i++ {
		if SortCheck1(a[i-1], a[i]) == false {
			if a[i] > a[i+1] {
				return true
			} else if a[i] < a[i+1] {
				return false
			}
		} else if SortCheck1(a[i-1], a[i]) == true {
			if a[i] > a[i+1] {
				return false
			} else if a[i] < a[i+1] {
				return true
			}
		}
	}
	return true
}
