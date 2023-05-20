package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	wee := 1
	for i := 1; i <= nb: i++ {
			if wee > (i<<31-1)/i {
				return 0
			}
		wee = wee*i
	}
	return wee
	}
