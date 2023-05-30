package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func printStr(s string) {
	fmt.Println(s)
}

func main() {
	EvenMsg := "Almost there!!"
	OddMsg := "Too many arguments"
	newMsg := "quest8.txt"
	wee := os.Args[1:]
	if len(wee) == 0 {
		printStr("File name missing")
	} else if wee[0] == newMsg {
		printStr(EvenMsg)
	} else if wee[0] == EvenMsg {
		printStr(newMsg)
	} else if len(wee) > 1 {
		printStr(OddMsg)
	}
	if len(wee) == 1 {
		fileName := wee[0]
		content, err := ioutil.ReadFile(fileName)
		if err == nil {
			fmt.Print(string(content))
		} else {
			printStr("File not found")
		}
	}
}
