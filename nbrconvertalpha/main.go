package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	caps := false
	for _, v := range arguments {
		if v == "--upper" {
			caps = true
		}
	}
	for _, arg := range arguments {
		numv := 0
		for _, v := range arg {
			numv = numv*10 + int(v-'0')
		}
		if numv >= 1 && numv <= 26 {
			if caps == false {
				z01.PrintRune(rune(numv - 1 + 'a'))
			} else {
				z01.PrintRune(rune(numv - 1 + 'A'))
			}
		}
	}
	z01.PrintRune('\n')
}
