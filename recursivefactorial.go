package piscine

func RecursiveFactorial(nb int) int {
	result := 1
	if nb < 0 || nb > 21 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	if nb <= 25 && nb > 0 {
		if nb == 1 {
			return nb
		}
		if nb == 2 {
			return nb
		}
		if nb == 3 {
			return 6
		}
		if nb == 4 {
			return 24
		}
		if nb == 5 {
			return 120
		}
		if nb == 6 {
			return 720
		}
		if nb == 7 {
			return 5040
		}
		if nb == 8 {
			return 40320
		}
		if nb == 9 {
			return (9 * 8 * 7 * 6 * 5 * 4 * 3 * 2)
		}
		if nb == 10 {
			return (10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2)
		}
		if nb == 11 {
			return (11 * 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2)
		}
		if nb == 12 {
			return (12 * 11 * 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2)
		}
		if nb == 13 {
			return (13 * 12 * 11 * 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2)
		}
		if nb == 14 {
			return (14 * 13 * 12 * 11 * 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2)
		}
		if nb == 15 {
			return (15 * 14 * 13 * 12 * 11 * 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2)
		}
		if nb == 16 {
			return (16 * 15 * 14 * 13 * 12 * 11 * 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2)
		}
		if nb == 17 {
			return (17 * 16 * 15 * 14 * 13 * 12 * 11 * 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2)
		}
		if nb == 18 {
			return (18 * 17 * 16 * 15 * 14 * 13 * 12 * 11 * 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2)
		}
		if nb == 19 {
			return (19 * 18 * 17 * 16 * 15 * 14 * 13 * 12 * 11 * 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2)
		}
		if nb == 20 {
			return (20 * 19 * 18 * 17 * 16 * 15 * 14 * 13 * 12 * 11 * 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2)
		}
	}
	return result
}
