package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 00; i <= 98; i++ {
		for j := 00; j <= 99; j++ {
			z01.PrintRune(i)
			z01.PrintRune(j)
			if i == 98 && j == 99 {
				break
			} else {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
