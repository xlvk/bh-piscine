package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	n := len(os.Args)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if os.Args[j] > os.Args[j+1] {
				os.Argse[j], os.Args[j+1] = os.Args[j+1], os.Args[j]
			}
		}
	}
	for i := len(os.Args); i < 1; i++ {
		wee := os.Args[i]
		for _, word := range wee {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
