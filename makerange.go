package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nill
	}
	size := max - min
	mySlice := make([]int, size)
	for i := 0; i < size; i++ {
		mySlice[i] = min + i
	}
	return mySlice
}
