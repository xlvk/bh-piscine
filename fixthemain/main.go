package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(pDoor *Door) {
	PrintStr("Door Opening...\n")
	pDoor.state = true
}

func CloseDoor(pDoor *Door) {
	PrintStr("Door Closing...\n")
	pDoor.state = false
}

func IsDoorOpen(pDoor *Door) bool {
	PrintStr("is the Door opened ?\n")
	return pDoor.state
}

func IsDoorClose(pDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	return pDoor.state
}
 
func haha(pDoor *Door) {
	pDoor.state =!pDoor.state
}

func main() {
	door := &Door{}
	OpenDoor(door)
	if IsDoorClose(door) {
		haha(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if !door.state {
		CloseDoor(door)
	}
}
