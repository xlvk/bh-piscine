package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimal := 0
	// Convert the number from baseFrom to decimal
	for i := 0; i < len(nbr); i++ {
		decimal = decimal*len(baseFrom) + int(nbr[i]-'0')
	}
	// Convert the decimal number to baseTo
	result := ""
	for decimal > 0 {
		remainder := decimal % len(baseTo)
		result = string(baseTo[remainder]) + result
		decimal = decimal / len(baseTo)
	}

	return result
}