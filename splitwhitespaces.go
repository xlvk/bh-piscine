package piscine

func SplitWhiteSpaces(s string) []string {
	var wee []string
	var o string
	index := 0
	mm := 0
	for i := 0; i < len(s); i++ {
		if s == " " || i == len(s)-1 {
			if i == len(s)-1 {
				o = s[index : i+1]
				mm++
			} else {
				o = s[index:i]
				mm++ 
			}
			for k : = 0; k < mm; k++ {
				wee = append(wee, o)
				o = ""
				index = i + 1
			}
		}
	}
	return wee
}
