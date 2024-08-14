package functions

import "sort"

//returns the total number of steps the ants have to travel to reach the end plus the total length of all paths
func TestAllPaths(paths [][]string, ants int) (int, int) {

	sort.SliceStable(paths, func(i, j int) bool {

		return len(paths[i]) < len(paths[j])
	})

	i := 0
	pathsCount := make([]int, len(paths))
	for i, r := range paths {
		pathsCount[i] = len(r) - 2
	}

	for ants != 0 {
		pathsCount[i] = pathsCount[i] + 1
		ants--
		if i == len(paths)-1 && pathsCount[i] > pathsCount[0] {
			i = 0
		} else if i != len(paths)-1 && pathsCount[i] > pathsCount[i+1] {
			i = i + 1
		}
	}

	total := 0
	roomNbr := 0
	for in, n := range pathsCount {
		if i != len(paths[in]) {
			if n > total {
				total = n
			}
		}
		roomNbr = roomNbr + len(paths[in])

	}
	return total, roomNbr
}
