package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	var sliceOfArgs []string
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 { // checks if something has been writen to standard output
		scanner := bufio.NewScanner(os.Stdin) //scans the standard input
		scanner.Split(bufio.ScanLines)        //splits standard input by newline
		for scanner.Scan() {
			sliceOfArgs = append(sliceOfArgs, scanner.Text()) //adds the resulting text of the scan to a string slice
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Error")
		return
	}
	stringInts := os.Args[1]
	stringSplit := strings.Split(stringInts, " ")
	sliceIntsA := []int{}
	for _, r := range stringSplit {
		temp, _ := strconv.Atoi(r)
		sliceIntsA = append(sliceIntsA, temp)
	}
	sliceIntsB := []int{}
	for _, str := range sliceOfArgs {
		if str == "pa" {
			//push the top first element of stack b to stack a
			temp := []int{sliceIntsB[0]}
			sliceIntsA = append(temp, sliceIntsA...)
			sliceIntsB = sliceIntsB[1:]
		}
		if str == "pb" {
			//push the top first element of stack a to stack b
			temp := []int{sliceIntsA[0]}
			sliceIntsB = append(temp, sliceIntsB...)
			sliceIntsA = sliceIntsA[1:]
		}
		if str == "sa" {
			//swap first 2 elements of stack a
			sliceIntsA[0], sliceIntsA[1] = sliceIntsA[1], sliceIntsA[0]
		}
		if str == "sb" {
			//swap first 2 elements of stack b
			sliceIntsB[0], sliceIntsB[1] = sliceIntsB[1], sliceIntsB[0]
		}
		if str == "ss" {
			//execute sa and sb
			sliceIntsA[0], sliceIntsA[1] = sliceIntsA[1], sliceIntsA[0]
			sliceIntsB[0], sliceIntsB[1] = sliceIntsB[1], sliceIntsB[0]
		}
		if str == "ra" {
			// rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
			sliceIntsA = append(sliceIntsA, sliceIntsA[0])
			sliceIntsA = sliceIntsA[1:]
		}
		if str == "rb" {
			//rotate stack b
			sliceIntsB = append(sliceIntsB, sliceIntsB[0])
			sliceIntsB = sliceIntsB[1:]
		}
		if str == "rr" {
			//execute ra and rb'
			sliceIntsA = append(sliceIntsA, sliceIntsA[0])
			sliceIntsB = append(sliceIntsB, sliceIntsB[0])
			sliceIntsB = sliceIntsB[1:]
			sliceIntsA = sliceIntsA[1:]
		}
		if str == "rra" {
			//reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
			temp := []int{sliceIntsA[len(sliceIntsA)-1]}
			sliceIntsA = append(temp, sliceIntsA...)
			sliceIntsA = sliceIntsA[:len(sliceIntsA)-1]
		}
		if str == "rrb" {
			//reverse rotate b
			temp := []int{sliceIntsB[len(sliceIntsB)-1]}
			sliceIntsB = append(temp, sliceIntsB...)
			sliceIntsB = sliceIntsB[:len(sliceIntsB)-1]
		}
		if str == "rrr" {
			//execute rra and rrb
			temp := []int{sliceIntsA[len(sliceIntsA)-1]}
			sliceIntsA = append(temp, sliceIntsA...)
			sliceIntsA = sliceIntsA[:len(sliceIntsA)-1]
			temp = []int{sliceIntsB[len(sliceIntsB)-1]}
			sliceIntsB = append(temp, sliceIntsB...)
			sliceIntsB = sliceIntsB[:len(sliceIntsB)-1]
		}
	}
	if len(sliceIntsB) == 0 && sort.IntsAreSorted(sliceIntsA) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
		//fmt.Println(sliceIntsA)
	}
}
