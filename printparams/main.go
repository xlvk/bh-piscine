package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		wee := os.Args[i]
		for _, word := range wee {
			z01.PrintRune(word)
		}
		z01.PrintRune('/n')
	}
}
