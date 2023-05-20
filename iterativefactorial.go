package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb > 25 {
		return 0
	}
	if nb <= 0 {
		return 1
	}
	if nb <= 25 && nb > 0 {
		for i := 1; i <= nb; i++ {
			result = result * i
		}
	}
	return result
}
