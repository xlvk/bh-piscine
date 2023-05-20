package piscine

import "github.com/01-edu/z01"

func IterativeFactorial(nb int) int {
if nb < 0 {
	return 0
}
wee := 1
for i := 1; i <= nb: i++ {
wee = wee*i
}
return wee
}
