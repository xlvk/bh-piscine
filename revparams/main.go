package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i > 1; i-- {
		wee := os.Args[i]
		for _, word := range wee {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
