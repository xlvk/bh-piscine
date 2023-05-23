package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	wee := os.Args[1:]
	runes := []rune(wee)
	for _, word := range runes {
			z01.PrintRune(word)
	}
}
