package main

import (
	"push-swap/library"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	A := library.FormatListOfInts(os.Args[1])
	if A == nil {
		return
	}
	B := []int{}
	instructions := strings.Split(os.Args[2], " ")
	library.ReadInstructionsAndExecute(A, B, instructions)
}
