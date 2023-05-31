package piscine

func DescendAppendRange(max, min int) []int {
	if max < min && (min != 0 || max != 0) {
		return []int{}
	} else if max == 0 && min == 0 {
		return nil
	}
	var myArray []int
	for i := max; i > min; i-- {
		myArray = append(myArray, i)
	}
	return myArray
}
