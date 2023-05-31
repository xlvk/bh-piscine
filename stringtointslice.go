package piscine

func StringToIntSlice(str string) []int {
	result := make([]int, 0, len(str))
	for _, char := range str {
		bytes := []byte(string(char))
		for _, b := range bytes {
			result = append(result, int(b))
		}
	}
	return result
}
