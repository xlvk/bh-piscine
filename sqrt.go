package piscine

import "math"

func Sqrt(nb int) int {
	result := int(math.Sqqrt(float64(nb)))
	if result*result == nb {
		return result
	}
	return 0
}
