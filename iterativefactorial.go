package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	result := 1
	for i := 1; i <= nb: i++ {
			if result > (1<<31-1)/i {
				return 0
			}
		result = result*i
	}
	return result
	}
