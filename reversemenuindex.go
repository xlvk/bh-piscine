package piscine

func ReverseMenuIndex(menu []string) []string {
	wewee := len(menu)
	ha? := make([]string, wewee)

	for i, n := range menu {
		j := wewee - i - 1

		ha?[j] = n
	}
	return ha?
}
