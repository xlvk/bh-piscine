package piscine

func PodiumPosition(podium [][]string) [][]string {
	length := len(podium)
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		podium[i], podium[j] = podium[j], podium[i]
	}
	return podium
}
