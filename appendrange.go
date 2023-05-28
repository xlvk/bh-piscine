package piscine

func AppendRange(min, max int) []int {
	size := max - min + 1
	var myArray []int
	for i := 0; i < size; i++ {
		myArray[i] = min + i
	}
	return myArray
}
