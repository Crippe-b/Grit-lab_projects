package main

import (
	"fmt"
	"lem-in/functions"
	"os"
)

//does things
func main() {
	if len(os.Args) == 2 {
		functions.Read(os.Args[1])
	} else {
		fmt.Println("Usage: go run . <filename.txt>")
	}
}
