package library

import (
	"fmt"
	"strings"
)

func ReadInstructionsAndExecute(A, B []int, instructions []string) {
	for i := 0; i < len(instructions); i++ {
		in := instructions[i]
		A, B = ExecuteCommand(A, B, strings.ToLower(in))
		fmt.Println(A, B, instructions[i])
		if Check(A) && len(B) == 0 && i != len(instructions)-1 {
			fmt.Println("OK but not finished")
		}
	}
	if Check(A) && len(B) == 0 {
		fmt.Println("OK")
	} else  {
		fmt.Println("KO")
	}
}
