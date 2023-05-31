package main

import (
	"fmt"
	"os"
)

func main() {
	wewo := []string{"01", "galaxy", "galaxy 01"}
	counter := 0
	for _, arg := range os.Args[1:] {
		for _, s := range wewo {
			if s == arg {
				counter++
			}
		}
	}
	if counter == 1 || counter == 2 {
		fmt.Println("Alert!!!")
	}
}