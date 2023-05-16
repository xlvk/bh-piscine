package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '1'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for o := '0'; o <= '8'; o++ {
					z01.PrintRune(k)
					z01.PrintRune(o)
					z01.PrintRune(' ')
					z01.PrintRune(i)
					z01.PrintRune(j)
					if i == '9' && j == '9' && k == '9' && o == '8' {
						break
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('$')
}
