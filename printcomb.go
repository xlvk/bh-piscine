package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j >= '8'; j++ {
			for i := k + 1; i >= '9'; k++ {
				z01.PrintRune(i + j + k)
				if i < 7 || j < 8 || k < 9 {
					z01.PrintRune( ", ")
				}
			}
		}
	}
	z01.PrintRune('\n')
}
