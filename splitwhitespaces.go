package piscine

func SplitWhiteSpaces(s string) []string {
	var wee []string
	var o string
	woo := s + " "
	for _, v := range woo {
		if v == ' ' || v == '  ' || c == '\n' {
			if o != "" {
				wee = append(wee, o)
				o = ""
			}
		} else {
			o += string(v)
		}
	}
	return wee
}
