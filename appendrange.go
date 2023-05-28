package piscine

func AppendRange(min, max int) []int {
	size := max - min + 1
	myArray := make([]int, size)
	for i := 0; i < size; i++ {
		myArray[i] = min + i
	}
	return myArray
}
