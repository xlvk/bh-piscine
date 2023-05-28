package piscine

func Split(s, sep string) []string {
	index := 0
	var o string
	var wee []string
	for i := 0; i < len(s); i++ {
		if s[i] == sep[0] && s[i+1] == sep[1] {
			if i == len(s)-1 {
				o = s[index : i+1]
			} else {
				o = s[index:i]
			}
			wee = append(wee, o)
			o = ""
			index = i + 2
		}
	}
	return wee
}
