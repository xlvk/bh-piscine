package main

import (
	// "errors"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	wee := os.Args[1:]
	if len(wee) == 0 {
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			printError(err)
			return
		}
		printRunes(input)
	} else {
		// first := true
		for _, arg := range wee {
			file, err := os.Open(arg)
			if err != nil {
				// err1 := errors.New("ERROR: open asd: no such file or directory")
				// s := err.Error()
				printError(err)
				continue
			}
			f, err := ioutil.ReadAll(file)
			if err != nil {
				printError(err)
				continue
			}
			// if !first {
			// 	z01.PrintRune('\n')
			// }
			printRunes(f)
			// first = false
			file.Close()
		}
	}
}

func printRunes(runes []byte) {
	for _, r := range string(runes) {
		z01.PrintRune(rune(r))
	}
}

func printError(err error) {
	errMsg := []rune("ERROR: " + err.Error())
	// errMsg1 := "ERROR: "
	// for i,ch := range errMsg1 {
	// 	z01.PrintRune(ch)
	// }
	for i := range errMsg {
		z01.PrintRune(errMsg[i])
	}
	z01.PrintRune('\n')
	os.Exit(1)
}
