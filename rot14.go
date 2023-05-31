package piscine

func Rot14(s string) string {
	resulte := ""
	for _, r := range s {
		if r == ' ' {
			resulte += " "
		} else if r >= 'A' && r <= 'L' || r >= 'a' && r <= 'l' {
			resulte += string(r + 14)
		} else if r >= 'M' && r <= 'Z' || r >= 'n' && r <= 'z' {
			resulte += string(r - 12)
		} else {
			resulte += string(r)
		}
	}
	return resulte
}
