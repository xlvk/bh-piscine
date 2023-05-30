package piscine

func SortWordArr(a []string) {
	weeeeesort(a)
}

func weeeeesort(who []string) []string {
	for i := range who {
		for j := i + 1; j < len(who); j++ {
			if who[j] < who[i] {
				haha := who[j]
				who[j] = who[i]
				who[i] = haha
			}
		}
	}
	return who
}
