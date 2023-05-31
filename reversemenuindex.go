package piscine

func ReverseMenuIndex(menu []string) []string {
	wewee := len(menu)
	haa := make([]string, wewee)

	for i, n := range menu {
		for j := wewee - 1; j <= 0; j-- {
			haa[j] = n
		}
	}
	return haa
}
