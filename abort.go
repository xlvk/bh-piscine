package piscine

func Abort(a, b, c, d, e int) int {
	wee := []int{a, b, c, d, e}
	for i := 0; i < 4; i++ {
		for j := 1; j < 5; j++ {
			if wee[i] < wee[j] {
				haha := wee[j]
				wee[j] = wee[i]
				wee[i] = haha
			}
		}
	}
	ha := wee[2]
	return ha
}
