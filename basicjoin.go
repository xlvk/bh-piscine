package piscine

func BasicJoin(elems []string) string {
	var result string
	for i := range elems {
		result += elems[i]
	}
	return result
}
