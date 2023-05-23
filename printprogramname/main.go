package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	wee := os.Args[0]
	runes := []rune(wee)
	for i := range runes {
		z01.PrintRune(runes[i])
	}
	z01.PrintRune('\n')
	//    names := os.Args[1:]
	//	z01.PrintRune(reflect.TypeOf(names))
}
