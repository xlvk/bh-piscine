package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	for _, ch := range n {
		z01.PrintRune(n)
	}
}
