package piscine

func PrintNbr(num int) {
}

func ForEach(f func(int), a []int) {
	for i, ch := range a {
		a = f(a[i])
	}
}
