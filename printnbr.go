package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
    digits := []byte{}
    negative := false
    if n < 0 {
        negative = true
        n = -n
    }
    for n > 0 {
        digit := n % 10
        digits = append([]byte{'0' + byte(digit)}, digits...)
        n /= 10
    }
    if negative {
        z01.PrintRune("-")
    }
    if len(digits) == 0 {
        z01.PrintRune("0")
    } else {
		z01.PrintRune(string(digits))
    }
}
