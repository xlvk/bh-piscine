package piscine

import "github.com/01-edu/z01"

func printInt(n int) {
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
        fmt.Print("-")
    }

    if len(digits) == 0 {
        fmt.Print("0")
    } else {
        fmt.Print(string(digits))
    }
}
