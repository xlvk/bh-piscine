package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}
	c := ""
	for i := 0; i < len(args); i++ {
		c += args[i]
		if i < len(args)-1 {
			c += "\n"
		}
	}
	return c
}
