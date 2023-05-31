package piscine

func ReverseMenuIndex(menu []string) []string {
	wewee := len(menu)
	haa := make([]string, wewee)

	for _, n := range menu {
		j := wewee - i - 1
		haa[j] = n
	}
	return haa
}
