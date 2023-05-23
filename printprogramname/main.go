package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	wee := os.Args[0]
	runes := []rune(wee)
	for _, word := range runes {
		if word == '.' || word == '/' {
		} else {
			z01.PrintRune(word)
		}
	}
	z01.PrintRune('\n')
}
