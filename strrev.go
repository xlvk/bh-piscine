package piscine

func StrRev(s string) string {
	runes := []rune(s)
	length := len(runes)
	for i, j:= '0', length-1; i < j; i, j = i + 1, j - 1 {
		rune[i], rune[j] = rune[j], rune[i] 
	}
}

