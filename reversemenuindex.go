package piscine

func ReverseMenuIndex(menu []string) []string {
	wewee := len(menu)
	haa := make([]string, wewee)

	for o, n := range menu {
		j := wewee - o - 1
		haa[j] = n
	}
	return haa
}
