package piscine

func StringToIntSlice(str string) []int {
	result := make([]int, len(str))
	for i, char := range str {
		result[i] = int(char)
	}
	return result
}
