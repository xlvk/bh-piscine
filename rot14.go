package piscine

func Rot14(s string) string {
	resulte := ""
	for _, r := range s {
		if r == ' ' {
			resulte += " "
		} else if r >= 'A' && r <= 'N' || r >= 'a' && r <= 'n' {
			resulte += string(r + 14)
		} else if r >= 'O' && r <= 'Z' || r >= 'o' && r <= 'z' {
			resulte += string(r - 12)
		} else {
			resulte += string(r)
		}
	}
	return resulte
}
