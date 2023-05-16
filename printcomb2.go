package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				if i == '0' && j == '0' && k == '0' {
					for o := '1'; o <= '9'; o++ {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(o)
						if !(i == '9' && j == '8' && k == '9' && o == '9') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				} else {
					for o := '0'; o <= '9'; o++ {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(o)
						if !(i == '9' && j == '8' && k == '9' && o == '9') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('$')
}
