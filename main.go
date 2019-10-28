package main

import (
	"fmt"
	"morsecode/morse"
)

func main() {
	fmt.Println(morse.Encode("hello world"))
	fmt.Println(morse.Decode(morse.Encode(".,?'!/()&:;=+-_\"$@")))
}
