package main

import (
	"os"

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
		valid := true
		var char rune

		if len(arg) != 1 {
			valid = false
		} else {
			char = rune(arg[0])
			if !('1' <= char && char <= '9') {
				valid = false
			}
		}

		if valid {
			if upperCase {
				char = char + ('A' - '1')
			} else {
				char = char + ('a' - '1')
			}
			z01.PrintRune(char)
		} else {
			z01.PrintRune(' ')
		}

		if i < len(args)-1 {
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')
}
