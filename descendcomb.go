package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '1'; j-- {
			for k := i; k >= '0'; k-- {
				for n := j - 1; n >= '0'; n-- {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(n)
					if i == '0' && j == '1' && k == '0' && n == '0' {
						break
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}
