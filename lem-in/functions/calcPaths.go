package functions

import (
	"fmt"
	"log"
)

//checks if slice are the same
func sliceSame(slice, compSlice []string) bool {
	if len(slice) != len(compSlice) {
		return false
	}
	for i, string1 := range slice {
		if string1 != compSlice[i] {
			return false
		}
	}
	return true
}

// checks if slice exists in slice of slice
func sliceExistsInSos(paths [][]string, compareSlice []string) bool {

	for _, slice := range paths {
		if sliceSame(slice, compareSlice) {
			return true
		}
	}
	return false
}

//checks if slice of slice are the same
func SosAreTheSame(paths, pathsCompare [][]string) bool {
	if len(paths) != len(pathsCompare) {
		return false
	}
	for _, pathCompare := range pathsCompare {
		if !sliceExistsInSos(paths, pathCompare) {
			return false
		}
	}
	return true
}

//checks if slice of slice already exists in slice of slice of slice
func soSAlreadyThere(pathPairs [][][]string, compSos [][]string) bool {
	for _, paths := range pathPairs {
		if SosAreTheSame(paths, compSos) {
			return true
		}
	}
	return false
}

//checks if paths does not share any rooms that aren't the start and end room
func isPossibleToCombine(paths [][]string, possiblePath []string, start, end string) bool {
	for _, compareRoom := range possiblePath {
		for _, path := range paths {
			for _, room := range path {
				if room == compareRoom && room != end && room != start {
					return false
				}
			}
		}
	}
	return true
}

//combines all possible paths that aren't duplicates and don't contain the same rooms (that aren't the start and end room)
func combinePaths(maxpaths int, allPaths [][]string, start, stop string) [][][]string {
	var firstPath []string
	firstPath = allPaths[0]
	for _, p := range allPaths {
		if len(p) < len(firstPath) {
			firstPath = p
		}
	}
	var pathPairs [][][]string
	var temp [][]string
	temp = append(temp, firstPath)
	pathPairs = append(pathPairs, temp)
	for i, path := range allPaths {
		var tempPair [][]string
		tempPair = append(tempPair, path)
		for i2, otherPath := range allPaths {
			if i != i2 {
				if isPossibleToCombine(tempPair, otherPath, start, stop) {
					tempPair = append(tempPair, otherPath)
				}
			}
		}
		if len(tempPair) > 1 && !soSAlreadyThere(pathPairs, tempPair) {
			pathPairs = append(pathPairs, tempPair)
		}
	}
	return pathPairs
}

//checks if slice contains provided string
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

//checks the sanity of the programmer while trying to get algoritms to work
func calculateValueOfRoute(route []string) int {
	counter := 0
	for i, room := range route {
		if i != len(route)-1 && Level[room].weight == -2 && Level[route[i+1]].weight == -2 {
			counter--
		} else {
			counter++
		}
	}
	return counter
}

//function that finds all possible paths that leads to end
func findShortest2(start, end string) [][]string {
	paths := [][]string{}
	paths = append(paths, []string{start})
	addedSomething := true
	hasReachedEndOnce := false
	for addedSomething {
		addedSomething = false
	loop:
		tempRoutes := [][]string{}
		for i, route := range paths {
			currentRoom := route[len(route)-1]
			if currentRoom == end {
				hasReachedEndOnce = true
			}
			if len(Level[currentRoom].connections) > 0 && currentRoom != end {
				for _, connection := range Level[currentRoom].connections {
					if !contains(route, connection) {
						temp := make([]string, 0)
						temp = append(temp, route...)
						temp = append(temp, connection)
						paths = append(paths, temp)
						addedSomething = true
						tempRoutes = append(tempRoutes, append(route, connection))
						val := calculateValueOfRoute(temp)
						if val <= Level[connection].value {
							if val == Level[connection].value {
								if len(temp) < len(Level[connection].lowestPath) {
									Level[connection].value = val
									Level[connection].lowestPath = temp
								}
							} else {
								Level[connection].value = val
								Level[connection].lowestPath = temp
							}

						}
					}
				}
				if len(tempRoutes) > 0 || (!addedSomething && currentRoom != end) {
					paths = deleteSsByIndex(paths, i)
					goto loop
				}
			}
		}
	}
	if !hasReachedEndOnce {
		log.Fatal("ERROR: no room reaches end")
	}

	return paths
}

//deletes slice of slice element, determined by index provided
func deleteSsByIndex(slice [][]string, i int) [][]string {
	copy(slice[i:], slice[i+1:])
	slice = slice[:len(slice)-1]
	return slice
}
func CalcPaths(start, end string, pissmyror int) {

	counter := 0
	_ = counter
	if len(Level[start].connections) > len(Level[end].connections) {
		counter = len(Level[end].connections)
	} else {
		counter = len(Level[start].connections)
	}
	allPaths := make([][]string, 0)
	allPaths = findShortest2(start, end)

	allPathsNumbered := combinePaths(counter, allPaths, start, end)

	var pathsCount []int
	var roomCount []int
	for _, paths := range allPathsNumbered {

		steps, roomNbr := TestAllPaths(paths, pissmyror)
		pathsCount = append(pathsCount, steps)
		roomCount = append(roomCount, roomNbr)
	}

	for _, paht := range allPathsNumbered {

		for _, p := range paht {
			if p[0] != start {
				log.Fatal("fuck", p[0])
			}
		}
	}
	i := 0
	for in, steps := range pathsCount {
		if pathsCount[i] > steps {
			i = in
		}
		if pathsCount[i] == steps && roomCount[i] > roomCount[in] {
			i = in
		}
	}
	fmt.Println("final", allPathsNumbered[i])
	MoveAnts(allPathsNumbered[i], pissmyror)
}
