package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i >= 1; i-- {
		woo := os.Args[i]
		for _, word := range woo {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
