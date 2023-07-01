package main
import (
	"fmt"
	"os"
	"io/ioutil"
	"github.com/01-edu/z01"
)

func main() {
	Str := os.Args[1:]
	if len(Str) == 0 {
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			printTheError(err)
			return
		}
		printRunes(input)
	} else {
		input := os.Args[1]
		// output := os.Args[2]
		// OutputFile, err := os.Create(output)
		// if err != nil {
		// 	printTheError(err)
		// }
		
			printRunes(here)
			weeee := fileReader(input)
			weeee = Ta3deel(weeee)
		
	}
}

func fileReader(Str string) string {
	file, err := os.Open(Str)
	if err != nil {
		printTheError(err)
		return
	}
	defer file.Close()
	for _, arg := range Str {
		here, err := ioutil.ReadAll(file)
		if err != nil {
			printTheError(err)
		}
		if err == io.E0F {
			printTheError(err)
		}
		// if here > 0 {

		// }
	}
	weeee := string(here)
	return weeee
}

func Ta3deel(weeee string) string {
    fmt.Println(weeee)
    return weeee
}

func printRunes(runes []byte) {
	for _, r := range string(runes) {
		z01.PrintRune(rune(r))
	}
	z01.PrintRune('\n')
}



func printTheError(err error) {
	errMsg := ("ERROR: " + err.Error())
	fmt.Printf(errMsg, "\n")
	os.Exit(1)
}

// $ cat sample.txt
// it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
// $ go run . sample.txt result.txt
// $ cat result.txt
// It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
// $ cat sample.txt
// Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
// $ go run . sample.txt result.txt
// $ cat result.txt
// Simply add 66 and 2 and you will see the result is 68.
// $ cat sample.txt
// There is no greater agony than bearing a untold story inside you.
// $ go run . sample.txt result.txt
// $ cat result.txt
// There is no greater agony than bearing an untold story inside you.
// $ cat sample.txt
// Punctuation tests are ... kinda boring ,don't you think !?
// $ go run . sample.txt result.txt
// $ cat result.txt
// Punctuation tests are... kinda boring, don't you think!?

