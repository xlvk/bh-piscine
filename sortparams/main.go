package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	n := len(os.Args)
	for i := 1; i < n-1; i++ {
		for j := 1; j <= n-i-1; j++ {
			if os.Args[j] > os.Args[j+1] {
				os.Args[j], os.Args[j+1] = os.Args[j+1], os.Args[j]
			}
		}
	}
	for i := 1; i < len(os.Args); i++ {
		wee := os.Args[i]
		for _, word := range wee {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
