package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := j; k <= '9'; k++ {
				for o := '0'; o <= '9'; o++ {
					count := '0'
					count++
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					if count == 1 {
						z01.PrintRune(o + 1)
					} else {
						z01.PrintRune(o)
					}
					if i == '9' && j == '8' && k == '9' && o == '9' {
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
