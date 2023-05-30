package piscine

func PrintNbro(num int) {
}

func ForEach(f func(int), a []int) {
	for _, i := range a {
		f(i)
	}
}
