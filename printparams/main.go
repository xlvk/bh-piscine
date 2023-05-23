package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	wee := os.Args[1:]
	for _, word := range wee {
		z01.PrintRune(word)
	}
}
