package piscine

func CollatzCountdown(start int) int {
	wow := start
	if start <= 0 {
		return -1
	} else {
		if start%2 == 0 {
			return wow / 2
		} else {
			return (3 * wow) + 1
		}
	}
}
