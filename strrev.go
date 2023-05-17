package piscine

func StrRev(s string) string {
	v := []rune(s)
	length := len(v)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}
}
