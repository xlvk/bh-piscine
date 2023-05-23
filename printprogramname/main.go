package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	wee := os.Args[0]
	runes := []rune(wee)
<<<<<<< HEAD
	for _, word := range runes {
		if word == '.' || word == '/' {
		} else {
			z01.PrintRune(word)
=======
	for i, word := range runes {
		if i != 0 && i != 1 {
			z01.PrintRune(rune(word))
>>>>>>> 3bc26dd (Q.6, pt.1)
		}
	}
	z01.PrintRune('\n')
}
