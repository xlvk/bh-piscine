package piscine

func MakeRange(min, max int) []int {
	if min > max {
		return nil
	}
	size := max - min
	myArray := make([]int, size+1)
	for i := min; i < max; i++ {
		myArray[i] = min + i
	}

}
