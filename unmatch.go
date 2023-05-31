package piscine

func Unmatch(arr []int) int {
	for i, vo := range arr {
		we := 1
		for j, wa := range arr {
			if vo == wa && i != j {
				we++
			}
		}
		if we%2 == 1 {
			return vo
		}
	}
	return -1
}
