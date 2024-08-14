package main

import (
	"asciiGo/justify"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 4 {
		words := os.Args[1]
		asciiFont := os.Args[2]
		justifyInp := os.Args[3]
		justify.AsciiJustify(words, asciiFont, justifyInp)
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\n\nExample: go run . something standard --align=right")
		return
	}
}
