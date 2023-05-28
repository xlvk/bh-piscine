func AppendRange(min, max int) []int {
	size := max - min + 1
	var myArray []int
	for i := 0; i < size; i++ {
        myArray = append(myArray, min+i)
    }
	return myArray
}
