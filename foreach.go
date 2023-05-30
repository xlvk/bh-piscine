package piscine

func PrintNbr()
func ForEach(f func(int), a []int) {
	a = f(a)
	return a
}
