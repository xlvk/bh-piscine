package piscine

func StringToIntSlice(str string) []int {
	result := make([]int, len(str))
	for i, char := range str {
		if int(char) == 0 {
			continue
		} else {
			result[i] = int(char)
		}
	}
	return result
}
