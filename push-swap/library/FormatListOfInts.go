package library

import (
	"strconv"
	"strings"
	"fmt"
)

func FormatListOfInts(s string) []int {
	var A []int
	list := []string(strings.Split(s, " "))
	for _, nr := range list {
		if n, err := strconv.Atoi(nr); err == nil {
			A = append(A, n)
		} else if err != nil {
			fmt.Println("Error")
			return nil
		}
	}
	return A
}
