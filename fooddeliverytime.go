package piscine

type food struct {
	preptime int
}
  
func FoodDeliveryTime(order string) int {
	burger := "burger"
	chips := "chips"
	nuggets := "nuggets"
	if order == burger {
		return 15
	}
	if order == chips {
		return 10
	}
	if order == nuggets {
		return 12
	} else {
		return 404
	}
}
  