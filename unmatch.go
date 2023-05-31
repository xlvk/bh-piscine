package piscine

func Unmatch(a []int) int {
	me := 0
	if len(a) == 0 {
		return me
	}
	if len(a) == 1 {
		return a[0]
	}
	if len(a) == 2 {
		if a[0] != a[1] {
			return me
		}
		return a[0]
	}
	if len(a) == 3 {
		if a[0] != a[1] || a[0] != a[2] {
			return me
		}
		return a[0]
	}
	if len(a) == 4 {
		if a[0] != a[1] || a[0] != a[2] || a[0] != a[3] {
			return me
		}
		return a[0]
	}
	if len(a) == 5 {
		if a[0] != a[1] || a[0] != a[2] || a[0] != a[3] || a[0] != a[4] {
			return me
		}
		return a[0]
	}
	return me
}
