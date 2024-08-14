package library


func deleteElem(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])	// removes the element at index i
	slice = slice[:len(slice)-1]
	return slice
}

func shortenOp(operations []string) []string {
Loop:
	for i := 0; i < len(operations); i++ { // check if we can use rr or rrr instead of ra/rb and rra/rrb
		if operations[i] == "ra" || operations[i] == "rra" {
			i2 := 1
			stop := false
			for !stop {
				if operations[i] == "ra" && operations[i+i2] == "rb" {
					operations[i] = "rr"
					operations = deleteElem(operations, i+i2)
					goto Loop
				} else if operations[i] == "rra" && operations[i+i2] == "rrb" {
					operations[i] = "rrr"
					operations = deleteElem(operations, i+i2)
					goto Loop
				} else if operations[i+i2] == "rr" || operations[i+i2] == "rrr" || operations[i+i2] == "sa" {
					i2++
				} else {
					stop = true
				}
			}
		}
	}
	return operations
}

func getNum(slcIntA []int, startN int, stopN int, foundNum bool) (int, bool, bool) {
	counterfor1 := 0
	counterfor2 := 1
	for i := 0; i <= len(slcIntA)-1; i++ {
		if slcIntA[i] >= startN && slcIntA[i] <= stopN {
			foundNum = true
			i = len(slcIntA)
		} else {
			counterfor1++
		}
	}
	for i := len(slcIntA) - 1; i >= 0; i-- {
		if slcIntA[i] >= startN && slcIntA[i] <= stopN {
			foundNum = true
			i = 0
		} else {
			counterfor2++
		}
	}
	if counterfor2 < counterfor1 && foundNum{
		return counterfor2, false, foundNum
	} else if foundNum {
		return counterfor1, true, foundNum
	} else {
		return counterfor1, true, foundNum
	}
}

func fastestRotate(num int, slcIntB []int, opARot string, stepsA int) (int, string) {
	normalC := 0
	reverseC := 1
// find the fastest way to rotate stack a
	for i := 0; i < len(slcIntB)-1; i++ {	// normal rotation if its closer to the top
		if num == slcIntB[i] {
			i = len(slcIntB)
		} else {
			normalC++
		}
	}
	for i := len(slcIntB) - 1; i > 0; i-- {	// reverse rotation if its closer to bottom
		if num == slcIntB[i] {
			i = 0
		} else {
			reverseC++
		}
	}
	BRotRb := normalC
	BRotRRb := reverseC

	if opARot == "ra" {
		if BRotRb-stepsA < 0 {
			BRotRb = 0
		} else {
			BRotRb = BRotRb - stepsA
		}
	}
	if opARot == "rra" {
		if BRotRRb-stepsA < 0 {
			BRotRRb = 0
		} else {
			BRotRRb = BRotRRb - stepsA
		}
	}

	if BRotRb <= BRotRRb {
		return normalC, "rb"
	}
	return reverseC, "rrb"
}

func compB(num int, slcIntB []int, opARot string, stepsA int) (int, string) {
	isLarger := true
	isSmaller := true
	largestB := slcIntB[0]
	smallestB := slcIntB[0]
	for _, n := range slcIntB {	// find the largest and smallest number in stack b
		if n > largestB {
			largestB = n
		}
		if n < smallestB {
			smallestB = n
		}
		if n > num {
			isLarger = false
		}
		if n < num {
			isSmaller = false
		}
	}
	if (isLarger && !isSmaller) || (isSmaller && !isLarger) {
		steps, rotation := fastestRotate(largestB, slcIntB, opARot, stepsA)
		return steps, rotation
	} else {
		closestS := 0
		for _, n := range slcIntB {
			if closestS < n && num > n {
				closestS = n
			}
		}
		steps, rotation := fastestRotate(closestS, slcIntB, opARot, stepsA)
		return steps, rotation
	}

}

func rb(slcIntB []int, steps int, operations []string) ([]int, []string) {
	//rotate stack b
	for i := 0; i < steps; i++ {
		slcIntB = append(slcIntB, slcIntB[0])
		slcIntB = slcIntB[1:]
		operations = append(operations, "rb")
	}
	return slcIntB, operations
}

func rrb(slcIntB []int, steps int, operations []string) ([]int, []string) {
	//reverse rotate stack b
	for i := 0; i < steps; i++ {
		temp := []int{slcIntB[len(slcIntB)-1]}
		slcIntB = append(temp, slcIntB...)
		slcIntB = slcIntB[:len(slcIntB)-1]
		operations = append(operations, "rrb")
	}
	return slcIntB, operations
}

func putInB(slcIntA []int, slcIntB []int, operations []string, opARot string, stepsA int, startN int) ([]int, []int, []string) {
	num := slcIntA[0]
	if len(slcIntB) <= 1 {
		slcIntA, slcIntB = Pb(slcIntA, slcIntB)
		operations = append(operations, "pb")
		return slcIntA, slcIntB, operations
	}
	steps, rotation := compB(num, slcIntB, opARot, stepsA)
	if len(slcIntA) > 1 && (slcIntA[1] >= startN && slcIntA[1] <= startN+19) {
		num2 := slcIntA[1]
		steps2, rotation2 := compB(num2, slcIntB, opARot, stepsA)
		if steps-steps2 > 2 {
			slcIntA = Sa(slcIntA)
			operations = append(operations, "sa")
			steps = steps2
			rotation = rotation2
		}
	}
	if rotation == "rb" {
		slcIntB, operations = rb(slcIntB, steps, operations)
	} else {
		slcIntB, operations = rrb(slcIntB, steps, operations)
	}
	slcIntA, slcIntB = Pb(slcIntA, slcIntB)
	operations = append(operations, "pb")
	return slcIntA, slcIntB, operations
}

func runShit(slcIntA []int, operations []string, chunkNbr int) []string {
	slcIntB := []int{}
	numAddOrder := []int{}
	closestToTop := false
	stepsA := 0
	times := 0
	for len(slcIntA) > 0 {
		startN := times*chunkNbr
		stopN := startN + chunkNbr
		opARot := ""
		foundNum := false
		stepsA, closestToTop, foundNum = getNum(slcIntA, startN, stopN, foundNum)
		for i := 0; i < stepsA; i++ {
			if closestToTop && foundNum {
				slcIntA = Ra(slcIntA)
				operations = append(operations, "ra")
				opARot = "ra"
			} else if foundNum{
				slcIntA = Rra(slcIntA)
				operations = append(operations, "rra")
				opARot = "rra"
			}
		}
		if foundNum {
		numAddOrder = append(numAddOrder, slcIntA[0])
		slcIntA, slcIntB, operations = putInB(slcIntA, slcIntB, operations, opARot, stepsA, startN)
		} else if !foundNum {
			times++
		}
	}

	findLargestB := false
	largestB := 0
	for len(slcIntB) > 0 {
		stepsB := 0
		foundNum := false
		if !findLargestB {
			for _, n := range slcIntB {
				if n > largestB {
					largestB = n
				}
			}
			findLargestB = true
			stepsB, closestToTop, foundNum = getNum(slcIntB, largestB, largestB, foundNum)
			for i := 0; i < stepsB; i++ {
				if closestToTop {
					slcIntB = Rb(slcIntB)
					operations = append(operations, "rb")
				} else {
					slcIntB = Rrb(slcIntB)
					operations = append(operations, "rrb")
				}
			}

		}
		slcIntA, slcIntB = Pa(slcIntA, slcIntB)
		operations = append(operations, "pa")
	}


	operations = shortenOp(operations)
	return operations
}

func ManyNumbers(sliceInts []int) []string {
	operations := []string{}
	//sliceInts = []int{68, 1, 2, 63, 77, 90, 58, 79, 59, 4, 71, 31, 8, 44, 67, 48, 73, 5, 83, 54, 9, 39, 37, 82, 74, 60, 93, 28, 57, 15, 100, 12, 35, 13, 22, 78, 81, 76, 49, 52, 62, 27, 53, 32, 50, 42, 38, 96, 34, 17, 10, 89, 86, 16, 66, 56, 40, 25, 97, 41, 24, 92, 20, 69, 21, 30, 7, 64, 46, 23, 94, 70, 6, 36, 14, 11, 29, 18, 45, 91, 26, 75, 95, 55, 87, 61, 43, 3, 33, 99, 65, 51, 85, 47, 98, 84, 72, 80, 19, 88}
	//sliceInts = []int{45, 67, 55, 91, 62, 81, 65, 95, 27, 71, 75, 86, 19, 51, 85, 20, 98, 35, 29, 68, 11, 99, 73, 24, 36, 5, 46, 96, 44, 8, 70, 57, 39, 100, 43, 56, 26, 12, 82, 33, 84, 89, 32, 74, 40, 41, 25, 94, 4, 13, 2, 15, 54, 22, 83, 88, 63, 59, 3, 14, 6, 78, 48, 97, 61, 34, 77, 49, 52, 30, 79, 92, 28, 42, 17, 16, 53, 47, 87, 9, 31, 76, 38, 10, 7, 23, 64, 1, 72, 50, 60, 80, 37, 90, 21, 93, 69, 66, 18, 58}
	//sliceInts := []int{781, 37, 1888, 347, 8}
	//sliceInts := []int{2, 1, 3, 6, 5, 8}
	chunkNbr := 20
	counter := 0
	operations = runShit(sliceInts, operations, chunkNbr)
	if len(sliceInts) <= 100 {
		for len(operations) > 700 && counter < 200 {
			chunkNbr++
			counter++
			operations = []string{}
			operations = runShit(sliceInts, operations, chunkNbr)
		}
	}
	return operations

}
