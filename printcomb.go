package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for k := j + 1; k <= '9'; k++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				if i == '7' && k == '8' && j == '9' {
					break

				} else {
				z01.PrintRune(',')
				z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
