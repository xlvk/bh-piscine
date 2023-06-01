package main

import "github.com/01-edu/z01"

const OPEN = true
const CLOSE = false

type Door struct {
	state bool
}
func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(Door *Door) bool {
	PrintStr("Door Openning...\n")
	Door.state = OPEN
	return true
}

func CloseDoor(Door *Door) bool {
	PrintStr("Door Closing...\n")
	Door.state = CLOSE
	return true
}

func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?\n")
	return Door.state 
}

func IsDoorClose(Door *Door) bool {
	PrintStr("is the Door closed ?\n")
	return Door.state
}

func main() {
	door := Door{}
	if IsDoorClose(&door) {
		OpenDoor(&door)
	}
	if IsDoorOpen(&door) {
		CloseDoor(&door)
	}
	if door.state == OPEN {
		CloseDoor(&door)
	}
}
