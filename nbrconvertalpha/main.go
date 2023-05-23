package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upperCase := false

	if len(args) > 0 && args[0] == "--upper" {
		upperCase = true
		args = args[1:]
	}

	for i, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil || num < 1 || num > 26 {
			z01.PrintRune(' ')
		} else {
			if upperCase {
				z01.PrintRune(rune('A' + num - 1))
			} else {
				z01.PrintRune(rune('a' + num - 1))
			}
		}

		if i < len(args)-1 {
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')
}
