package library

func getNum2(slcIntA []int, startN int, stopN int, foundNum bool) (int, bool, bool) {
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

func SmallSort(a []int, b []int, max int, min int, steps []string) ([]int, []int, []string) {
	if len(a) == 6 {
		foundNum := false
		turnA := 0
		closestToTop := false

		for i := 0; i <= len(a) -1; i++ {
			if (a[i] != max && i != len(a) -1 && a[i] > a[i+1]) || (a[i] != max && i == len(a)-1 && a[i] < a[i-1]) {
				turnA, closestToTop, foundNum = getNum2(a, a[i], a[i], foundNum)
				for i2 := 0; i2 < turnA; i2++ {
					if closestToTop && foundNum {
						a = Ra(a)
						steps = append(steps, "ra")
					} else if foundNum {
						a = Rra(a)
						steps = append(steps, "rra")
					}
				}
				a = Sa(a)
				steps = append(steps, "sa")
				i = -1
			}
		}
		turnA, closestToTop, foundNum = getNum2(a, min, min, foundNum)
		for i3 := 0; i3 < turnA; i3++ {
			if closestToTop && foundNum {
				a = Ra(a)
				steps = append(steps, "ra")
			} else if foundNum {
				a = Rra(a)
				steps = append(steps, "rra")
			}
		}

		
		return a, b, steps
	}

	for len(a) > 3 {	//while stack a has more than 3 elements, push to b
		a, b = Pb(a, b)
		steps = append(steps, "pb")
	}

	switch {	//sort the 3 numbers in stack a by deciding which instruction to use based on order of numbers
	case a[0] > a[1] && a[1] < a[2] && a[2] > a[0]: //2 1 3
		a = Sa(a)
		steps = append(steps, "sa")
	case a[0] > a[1] && a[1] > a[2]: //3 2 1
		a = Sa(a)
		steps = append(steps, "sa")
		a = Rra(a)
		steps = append(steps, "rra")
	case a[0] > a[1] && a[1] < a[2] && a[0] > a[2]: //3 1 2
		a = Ra(a)
		steps = append(steps, "ra")
	case a[0] < a[1] && a[1] > a[2] && a[0] > a[2]: //2 3 1
		a = Rra(a)
		steps = append(steps, "rra")
	case a[0] < a[1] && a[1] > a[2] && a[2] > a[0]: //1 3 2
		a = Sa(a)
		steps = append(steps, "sa")
		a = Ra(a)
		steps = append(steps, "ra")
	default: //1 2 3
		break
	}

	var Index int
	for len(b) != 0 {	
		Index = FindClosestHigherIndex(a, b[0])	//finds the index of where to push back
		a, b, steps = FindWherePushBack(a, b, Index, steps) //pushes back to a based on index
	}
	minValueA, _ := FindMin(a)

	turns, shortestRa, foundNum := getNum(a, minValueA, minValueA, false)	//finds the shortest way to rotate stack a to get the min value to the top
	_ = foundNum

	for i := 0; i < turns; i++ { // if the bool shortestRa is true, ra, else rra
		if shortestRa {
			a = Ra(a)
			steps = append(steps, "ra")
		} else {
			a = Rra(a)
			steps = append(steps, "rra")
		}
	}
	return a, b, steps
}

func FindWherePushBack(a []int, b []int, Index int, steps []string) ([]int, []int, []string) {
	if len(a)-Index > Index {	//if the index is closer to the top of the stack, ra until it is at the top
		for i := 0; i < Index; i++ {
			a = Ra(a)
			steps = append(steps, "ra")
		}
		a, b = Pa(a, b)
		steps = append(steps, "pa")
	} else {	//if the index is closer to the bottom of the stack, rra until it is at the top
		for i := 0; i < len(a)-Index; i++ {
			a = Rra(a)
			steps = append(steps, "rra")
		}
		a, b = Pa(a, b)
		steps = append(steps, "pa")
	}
	return a, b, steps
}

func FindMin(x []int) (int, int) {
	minValue := x[0]
	var bIndex int
	for i := 0; i < len(x); i++ {	//compares all values through the stack and finds the min value
		if x[i] < minValue {
			minValue = x[i]
			bIndex = i
		}
	}
	return minValue, bIndex
}

func FindMax(x []int) (int, int) {
	maxValue := x[0]
	var index int
	for i := 0; i < len(x); i++ {	//compares all values through the stack and finds the max value
		if x[i] > maxValue {
			maxValue = x[i]
			index = i
		}
	}
	return maxValue, index
}

func FindClosestHigherIndex(x []int, value int) int {	// finds the index of where to push back
	_, minIndex := FindMin(x)
	max, _ := FindMax(x)
	for i := value + 1; i <= max; i++ {	
		for aIndex, n := range x {	//finds the index of the value that is closest to the value of b[0]
			if i == n {
				return aIndex
			}
		}
	}
	return minIndex
}

func Pa(a []int, b []int) ([]int, []int) {
	//push the top first element of stack b to stack a
	temp := []int{b[0]}
	a = append(temp, a...)
	b = b[1:]
	return a, b
}

func Pb(sliceIntsA []int, sliceIntsB []int) ([]int, []int) {
	//push the top first element of stack a to stack b
	temp := []int{sliceIntsA[0]}
	sliceIntsB = append(temp, sliceIntsB...)
	sliceIntsA = sliceIntsA[1:]
	return sliceIntsA, sliceIntsB
}

func Sa(a []int) []int {
	a[0], a[1] = a[1], a[0]
	return a
}

func Sb(b []int) []int {
	b[0], b[1] = b[1], b[0]
	return b
}

func Ss(a []int, b []int) ([]int, []int) {
	a = Sa(a)
	b = Sb(b)
	return a, b
}

func Ra(a []int) []int {
	a = append(a, a[0])
	a = a[1:]
	return a
}

func Rb(b []int) []int {
	b = append(b, b[0])
	b = b[1:]
	return b
}

func Rr(a []int, b []int) ([]int, []int) {
	a = Ra(a)
	b = Rb(b)
	return a, b
}

func Rra(a []int) []int {
	// reverse rotate stack a (shift down all elements of stack a by 1, the last element becomes the first one)
	//reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
	temp := []int{a[len(a)-1]}
	a = append(temp, a...)
	a = a[:len(a)-1]
	return a
}

func Rrb(b []int) []int {
	// reverse rotate stack b (shift down all elements of stack b by 1, the last element becomes the first one)
	temp := []int{b[len(b)-1]}
	b = append(temp, b...)
	b = b[:len(b)-1]
	return b
}

func Rrr(a []int, b []int) ([]int, []int) {
	a = Rra(a)
	b = Rrb(b)
	return a, b
}

func CheckIfDuplicateNumber(a []int) bool {
	//checks if there are duplicate numbers in the stack
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				return true
			}
		}
	}
	return false

	/*visited := make(map[int]bool, 0)
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] == true{
			return true
			} else {
				visited[arr[i]] = true
			}
		}
	return false */
}