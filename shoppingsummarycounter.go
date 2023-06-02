package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	var list []string
	var countList []string
	item := ""
	for _, runes := range str {
		if runes == ' ' {
			list = append(list, item)
			item = ""
		} else {
			item += string(runes)
		}
	}

	list = append(list, item)
	for _, i := range list {
		add := true
		for _, new := range countList {
			if new == i {
				add = false
			}
		}
		if add {
			countList = append(countList, i)
		}
	}

	shopList := make(map[string]int, len(countList))
	for _, i := range countList {
		shopList[i] = 0
	}

	for _, l := range list {
		shopList[l]++
	}
	return shopList
}
