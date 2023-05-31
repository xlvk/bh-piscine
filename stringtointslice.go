package piscine

func StringToIntSlice(str string) []int {
	result := make([]int, 0, len(str))
	for _, char := range str {
		result = append(result, int(char))
	}
	return result
}
