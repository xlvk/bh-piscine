package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0o0; i <= 98; i++ {
		for j := 0o0; j <= 99; j++ {
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
