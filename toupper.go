package piscine

func ToUpper(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			char -= 'a' - 'A'
		}
		result += string(char)
	}
	return result
}
