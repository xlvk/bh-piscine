package main

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune("T") 
	}
	if nb >= 0 {
		z01.PrintRune("F")
	}
}

func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}