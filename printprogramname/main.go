package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	wee := os.Args[0]
	for i := 2; i <= len(wee); i++ {
		z01.PrintRune(d)
		//    names := os.Args[1:]
		//	z01.PrintRune(reflect.TypeOf(names))
	}
	z01.PrintRune('\n')
}
