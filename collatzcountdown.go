package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	} else {
		if start%2 == 0 {
			return start / 2
		} else {
			return (3 * start) + 1
		}
	}
}
