package piscine

import "os"

func ShoppingSummaryCounter(str string) map[string]int {
	haha := make([]string, 0)
	mememe := ""
	d := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			mememe += string(str[d:i])
			haha = append(haha, mememe)
			mememe = ""
			d = i + 1
		} else if i == (len(str) - 1) {
			mememe += string(str[d : i+1])
			haha = append(haha, mememe)
		}
	}
	TheThing := os.Args
	resulte := make(map[string]int, 0)
	for j := 1; j < len(TheThing); j++ {
		count := 0
		for i := 0; i < len(haha); i++ {
			if haha[i] == TheThing[j] {
				count++
			}
		}
		resulte[TheThing[j]] = count
	}
	return resulte
}
