package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	caps := false

	if len(arguments) > 0 && arguments[0] == "--upper" {
		caps = true
		arguments = arguments[1:]
	}

	for _, arg := range arguments {
		for _, c := range arg {
			if c < '0' || c > '9' {
				z01.PrintRune(' ')
				break
			}
			numv := int(c - '0')
			if numv >= 1 && numv <= 26 {
				if caps {
					z01.PrintRune(rune(numv - 1 + 'A'))
				} else {
					z01.PrintRune(rune(numv - 1 + 'a'))
				}
			}
		}
	}
	z01.PrintRune('\n')
}
