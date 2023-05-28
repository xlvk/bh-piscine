package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, word := range a {
		PrintString(word)
		PrintRune('\n')
	}
}
