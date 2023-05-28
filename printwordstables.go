package piscine

func PrintWordsTables(a []string) {
	var words []string
	weee := len(a)
	word := ""
	for i := 0; i < weee; i++ {
		if a[i] != " " {
			word += string(a[i])
		} else if word != "" {
			words = append(words, word)
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)
	}
	a = words
	return a
}
