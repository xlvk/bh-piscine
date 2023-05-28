package piscine

func AppendRange(min, max int) []int {
	if min > max {
		return nil
	}
	var myArray []int
	for i := min; i <= max; i++ {
		myArray = append(myArray, i)
	}
	return myArray
}
