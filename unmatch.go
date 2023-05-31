package piscine

func Unmatch(a []int) int {
	me := -1
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}
	for i := 0; i < len(a); {
		for j := i + 1; j <= len(a)-1; j++ {
			if !checkMatch(a[i], a[j]) {
				me = a[i]
			}
		}
	}
	return me
}

func checkMatch(a, b int) bool {
	if a == b {
		return true
	}
	return false
}
