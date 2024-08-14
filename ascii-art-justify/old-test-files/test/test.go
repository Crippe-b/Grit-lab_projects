package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetTermDim1() {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	str := string(out[:len(out)-1])
	ind := 1
	for i, r := range str {
		if r == ' ' {
			ind = ind + i
		}
	}
	num, _ := strconv.Atoi(str[ind:])
	fmt.Println(num)

}

func GetTermDim2() string {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	size, err := cmd.Output()
	size = size[:len(size)-1]
	x := strings.SplitAfter(string(size), " ")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(x[1]))
	return string(x[1])
}

func main() {
	fmt.Println(GetTermDim2())
	GetTermDim1()
}
