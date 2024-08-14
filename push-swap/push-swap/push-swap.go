package main

import (
	"fmt"
	"os"
	"push-swap/library"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		return
	}

	stringInts := os.Args[1] // getting ints from command line
	stringSplit := strings.Split(stringInts, " ")
	sliceIntsA := []int{}
	for _, r := range stringSplit {
		temp, err := strconv.Atoi(r)
		if err != nil {
			fmt.Println("Error")
			return
		}
		sliceIntsA = append(sliceIntsA, temp)
	}
	sliceIntsB := []int{}
	maxValueA, _ := library.FindMax(sliceIntsA)
	minValueA, _ := library.FindMin(sliceIntsA)
	var steps []string

	if library.CheckIfDuplicateNumber(sliceIntsA) {
		fmt.Println("Error")
		return
	}
	if library.Check(sliceIntsA) { //if the stack is already sorted
		return
	} else if len(sliceIntsA) == 2 { //if the stack has only 2 elements
		if sliceIntsA[0] > sliceIntsA[1] {
			steps = append(steps, "sa")
			fmt.Println(steps[0])
		}
		return
	} else if len(sliceIntsA) <= 6 { //if the stack has 3, 4, 5 or 6 elements
		_, _, steps = library.SmallSort(sliceIntsA, sliceIntsB, maxValueA, minValueA, steps)
		//fmt.Println(steps) //print the steps
		for _, str := range steps {
			fmt.Println(str)
		}
	} else { //if the stack has more than 6 elements
		sliceOp := library.ManyNumbers(sliceIntsA)
		for _, str := range sliceOp {
			fmt.Println(str)
		}
		fmt.Println(len(sliceOp))
		return
	}
}
