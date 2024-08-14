package functions

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type AntFarm struct {
	ants        int
	connections []string
	lowestPath  []string
	weight      int
	value       int
	occupied    bool
	isStart     bool
	isEnd       bool
}

var Level = map[string]*AntFarm{}

//function that reads file and returns slice of strings
func Read(name string) {
	textLines := []string{}
	file, err := os.Open("audit/" + name)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		textLines = append(textLines, scan.Text())
	}
	file.Close()

	sliceToStruct(textLines)
}

func findName(str string) string {
	name := ""
	stop := false
	for _, r := range str {
		if r != ' ' && !stop {
			name = name + string(r)
		} else {
			stop = true
		}
	}
	return name
}

//function that finds the connections between rooms
func findConnections(str string) (string, string) {
	name1 := ""
	name2 := ""
	switchName := false
	for _, r := range str {
		if r == '-' {
			switchName = true
		}
		if !switchName {
			name1 = name1 + string(r)
		}
		if switchName && r != '-' {
			name2 = name2 + string(r)
		}
	}
	name1Exists := false
	name2Exists := false
	for name := range Level {
		if name == name1 {
			name1Exists = true
		}
		if name == name2 {
			name2Exists = true
		}
	}
	if !name1Exists || !name2Exists {
		log.Fatal("ERROR: tried to connect to a path that doesn't exist or used a '-' character for a room name")
	}
	return name1, name2
}

//function that sends our slices to the Antfarm struct
func sliceToStruct(slice []string) {
	startcheck := false
	endcheck := false
	start := ""
	end := ""
	pissmyra, err := strconv.Atoi(slice[0])
	if err != nil {
		log.Fatal("Error: either not a number or too large of a number")
	}
	slice = slice[1:]
	for i, str := range slice {
		if str[0] == 'L' {
			log.Fatal("Rooms can not start with the character 'L'")
		}
		if len(str) > 0 && str[0] != '#' && !strings.Contains(str, "-") {
			name := findName(str)
			s := new(AntFarm)
			s.ants = pissmyra
			s.isStart = true
			s.occupied = false
			s.weight = 0
			s.value = 1056
			s.lowestPath = []string{}
			Level[name] = s
		}
		if str == "##start" {
			startcheck = true
			name := findName(slice[i+1])
			s := new(AntFarm)
			s.isStart = true
			s.isEnd = false
			s.occupied = false
			s.weight = 0
			s.value = 1056
			s.lowestPath = []string{}
			Level[name] = s
			start = name
		}
		if str == "##end" {
			endcheck = true
			name := findName(slice[i+1])
			s := new(AntFarm)
			s.isEnd = true
			s.isStart = false
			s.occupied = false
			s.weight = 0
			s.value = 1056
			s.lowestPath = []string{}
			Level[name] = s
			end = name
		}
		if len(str) > 0 && str[0] != '#' && strings.Contains(str, "-") {
			name1, name2 := findConnections(str)
			Level[name1].connections = append(Level[name1].connections, name2)
			Level[name2].connections = append(Level[name2].connections, name1)
		}
	}

	//ERROR Handling
	if pissmyra < 1 || pissmyra > 100000 {
		log.Fatalf("ERROR: invalid number of Ants")
	}
	if !startcheck && !endcheck {
		log.Fatalf("OOPS read instructions bro, you need a start and an end room")
	}
	if !startcheck {
		log.Fatalf("ERROR: invalid data format, no start room found")
	}
	if !endcheck {
		log.Fatalf("ERROR: invalid data format, no end room found")
	}
	if len(Level[end].connections) == 0 {
		log.Fatalf("ERROR: no room reaches end")
	}

	CalcPaths(start, end, pissmyra)
}
