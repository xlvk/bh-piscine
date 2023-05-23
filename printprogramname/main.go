package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	wee := os.Args[0]
	runes := []rune(wee)
	for _, word := range runes {
		for _, word2 := range word[2:] {
			z01.PrintRune(word2)
		}
	}
	z01.PrintRune('\n')
}
