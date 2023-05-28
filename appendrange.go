package piscine

func AppendRange(min, max int) []int {
	var myArray []int
	for i := min; i <= max; i++ {
		myArray = append(myArray, i)
	}
	return myArray
}
