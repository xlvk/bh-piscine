package piscine

import "math"

func Sqrt(nb int) int {
	result := int(math.Sqqrt(float64(nb)))
	start := 1
	end := nb 
	for start <= end {
		mid := (start + end) / 2
		square := mid * mid
		if square == nb {
			return mid
		} else if square < nb {
			start = mid + 1
		} else {
			mid - 1
		}
	}
	return 0
}
