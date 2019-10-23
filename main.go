package main

import (
	"fmt"
	"morsecode/morse"
)

func main() {
	fmt.Println(morse.Encode("$", " ", "*"))
	fmt.Println(morse.Decode("...-..-", " ", "*"))
}
