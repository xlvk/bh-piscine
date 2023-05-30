package piscine

func PrintNbr(num int) {
}

func ForEach(f func(int), a []int) {
	for _, ch := range a {
		a = f(a)
	}
}
