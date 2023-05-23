package main

import ( 
	"os"
	"github.com/01-edu/z01"
 )

func main() {
	wee := os.Args
	z01.PrintRune(len(wee))
    names := os.Args[1:]
	z01.PrintRune(reflect.TypeOf(names))
}
