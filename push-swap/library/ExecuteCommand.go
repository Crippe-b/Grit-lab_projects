package library

import (
	"fmt"
)

func ExecuteCommand(a, b []int, command string) ([]int, []int) {
	if command == "pa" {
		a, b = Pa(a, b)
	} else if command == "pb" {
		a, b = Pb(a, b)
	} else if command == "sa" {
		a = Sa(a)
	} else if command == "sb" {
		b = Sb(b)
	} else if command == "ss" {
		a, b = Ss(a, b)
	} else if command == "ra" {
		a = Ra(a)
	} else if command == "rb" {
		b = Rb(b)
	} else if command == "rr" {
		a, b = Rr(a, b)
	} else if command == "rra" {
		a = Rra(a)
	} else if command == "rrb" {
		b = Rrb(b)
	} else if command == "rrr" {
		a, b = Rrr(a, b)
	} else {
		fmt.Println("Type one of the following instructions: [pa, pb, sa, sb, ss, ra, rb, rr, rra, rrb, rrr]")
	}
	return a, b
}

// func px(x, y []int) ([]int, []int) {
// 	if len(y) == 0 {
// 		fmt.Println("Error: Cannot push from empty stack")
// 		return x, y
// 	}
// 	return append([]int{y[0]}, x...), y[1:]
// }

// func sx(x []int) []int {
// 	if len(x) == 0 {
// 		fmt.Println("Error: Cannot swap elements in empty stack")
// 	} else {
// 		x[0], x[1] = x[1], x[0]
// 	}
// 	return x
// }

// func rx(x []int) []int {
// 	if len(x) == 0 {
// 		fmt.Println("Error: Cannot rotate elements in empty stack")
// 		return x
// 	}
// 	return append(x, x[0])[1:]
// }

// func rrx(x []int) []int {
// 	if len(x) == 0 {
// 		fmt.Println("Error: Cannot rotate elements in empty stack")
// 		return x
// 	}
// 	return append(x[len(x)-1:], x...)[:len(x)]
// }
