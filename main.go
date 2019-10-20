package main

import (
	"fmt"
	"morsecode/morse"
)

func main() {
	fmt.Println(morse.Encode(".,?'!/()&:;=+-_\"$@"))

	fmt.Println(morse.Decode(morse.Encode(".,?'!/()&:;=+-_\"$@")))
}
