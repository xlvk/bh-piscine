package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func main() {
	points := &point{}
	setPoint(points)
	PrintStr("x = ")
	PrintStr((points.x - '0'))
	PrintStr(", y = ")
	PrintStr((points.y - '0'))
	PrintStr("\n")
}
