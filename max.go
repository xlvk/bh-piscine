package piscine

func Max(a []int) int {
	leno := len(a)
	for i := 0; i < leno; i++ {
		for j := 0; j < leno; j++ {
			if a[j] > a[i] {
				haa := a[j]
				a[j] = a[i]
				a[i] = haa
			} else {
				continue
			}
		}
	}
	return a[len(a)-1]
}
