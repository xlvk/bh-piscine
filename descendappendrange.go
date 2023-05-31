package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		if min == 0 && max == 0 {
			return nil
		} else {
			return []int{}
		}
	}
	var myArray []int
	for i := max; i > min; i-- {
		myArray = append(myArray, i)
	}
	return myArray
}
