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
		if caps {
			z01.PrintRune(toUpper(arg))
		} else {
			z01.PrintRune(toLower(arg))
		}
	}
	z01.PrintRune('\n')
}

func toLower(arg string) rune {
	numv := 0
	for _, c := range arg {
		if c < '0' || c > '9' {
			return ' '
		}
		numv = numv*10 + int(c-'0')
	}
	if numv >= 1 && numv <= 26 {
		return rune(numv - 1 + 'a')
	}
	return ' '
}

func toUpper(arg string) rune {
	numv := 0
	for _, c := range arg {
		if c < '0' || c > '9' {
			return ' '
		}
		numv = numv*10 + int(c-'0')
	}
	if numv >= 1 && numv <= 26 {
		return rune(numv - 1 + 'A')
	}
	return ' '
}
