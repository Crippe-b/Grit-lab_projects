package functions

import (
	"fmt"
	"sort"
)

//Prints the ants movement according to shortest path
func MoveAnts(paths [][]string, pissmyror int) {

	sort.SliceStable(paths, func(i, j int) bool {

		return len(paths[i]) < len(paths[j])
	})
	pathsCount := make([]int, len(paths))
	for i, r := range paths {
		pathsCount[i] = len(r) - 2
	}
	antPaths := make([][]int, 0)
	for _, p := range paths {
		path := make([]int, len(p)-1)
		antPaths = append(antPaths, path)
	}

	pNbr := 0
	for pissmyror != 0 {
		pathsCount[pNbr] = pathsCount[pNbr] + 1
		pissmyror--
		if pNbr == len(paths)-1 && pathsCount[pNbr] > pathsCount[0] {
			pNbr = 0
		} else if pNbr != len(paths)-1 && pathsCount[pNbr] > pathsCount[pNbr+1] {
			pNbr = pNbr + 1
		}
	}
	antNbr := 0
	antsAtEnd := false

	for i := range pathsCount {
		pathsCount[i] = pathsCount[i] - len(paths[i]) + 2
	}

	steps := 0
	for !antsAtEnd {

		antsAtEnd = true

		for _, path := range antPaths {
			if len(path) == 1 {
				path[0] = 0
			}
			for i := len(path) - 2; i >= 0; i-- {
				if i+1 == len(path)-1 && path[i+1] != 0 {
					path[i+1] = 0
				}
				if i == len(path)-1 && path[i] != 0 {
					path[i] = 0
				}
				if path[i] != 0 {
					path[i+1] = path[i]
					path[i] = 0
					antsAtEnd = false

				}
			}
		}

		for i, pathNbr := range pathsCount {
			if pathNbr != 0 && antPaths[i][0] == 0 {
				antNbr++
				antPaths[i][0] = antNbr
				antsAtEnd = false
				pathsCount[i] = pathsCount[i] - 1
			}
		}
		for i2, antPath := range antPaths {

			for i3 := len(antPath) - 1; i3 >= 0; i3-- {
				if antPath[i3] != 0 {
					fmt.Print("L", antPath[i3], "-", paths[i2][i3+1], " ")
				}
			}

		}
		fmt.Println()
		steps++
	}
	fmt.Println("total steps:", steps-1)
}
