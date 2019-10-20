package main

import (
	"fmt"
	"morsecode/morse"
)

func main() {
	fmt.Println(morse.Encode("The quick brown fox jumps over 13 lazy dogs"))

	fmt.Println(morse.Decode("- .... . / --.- ..- .. -.-. -.- / -... .-. --- .-- -. / ..-. --- -..- / .--- ..- -- .--. ... / --- ...- . .-. / .---- ...-- / .-.. .- --.. -.-- / -.. --- --. ... "))
}
