package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		wee := os.Args[i]
		z01.PrintRune(word)
	}
}
