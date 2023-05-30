package piscine

func PrintNbr(num int) {
}

func ForEach(f func(int), a []int) {
	for i := range len(a)-1 {
		a = f(a[i])
	}
}
