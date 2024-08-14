package justify

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

func justify(result []string, termLength int) []string {
	for i := 0; i <= 7; i++ {
		for i2 := 0; i2 < termLength; i2++ {
			result[i] = result[i] + " "
		}
	}
	return result
}

func getTotalLength(words string, asciiFont string, txtlines []string) int {
	result := 0
	lenOfAscii := 0
	/*for _, r := range words {
		if r == ' ' && fontTax > 0 {
			fontTax--
		}
		switch r {
		case '!', '\'', ',', '.', '1', ':', ';', '`', 'i', 'l', '|' : result = result + 4
		case '(', ')', '^' : result = result + 5
		case ' ', '"', '$', '<', '>', '[', ']', 'f', 'j', 's', 't', 'z', '{', '}', '~' : result = result + 6
		case '?','K', 'c', 'e', 'k', 'r', 'x' : result = result + 7
		case '&', '-', '4', '=', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'L', 'M', 'O', 'P', 'Q', 'R', 'S', 'U', '_' : result = result + 9
		case '*', '@', 'T', 'Y' : result = result + 10
		case '#', 'A', 'V', 'w' : result = result + 11
		case 'm' : result = result + 12
		case 'W' : result = result + 15
		default : result = result + 9
		}
	}
	*/
	for _, r := range words {
		lenOfAscii = len(txtlines[int(r-' ')*9+2])
		result = result + lenOfAscii
	}
	return result
}

func getTermLength() int {
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
	return num

}

func checkAscii(str string) bool {
	for _, r := range str {
		if !(r >= ' ' && r <= '~') {
			return false
		}
	}
	return true
}

func printAscii(res []string, alignment string, termLength int) {
	// Justify string output in
	if res[0] == "" {
		fmt.Println()
		return
	}
	if alignment == "right" {
		termLength = termLength - len(res[0])
		for i := 0; i <= 7; i++ {
			for i2 := 0; i2 < termLength; i2++ {
				res[i] = " " + res[i]
			}
		}
	} else if alignment == "center" {
		termLength = termLength - len(res[0])
		termLength = termLength / 2
		for i := 0; i <= 7; i++ {
			for i2 := 0; i2 < termLength; i2++ {
				res[i] = " " + res[i]
			}
		}
	}
	for i := 0; i <= 7; i++ {

		fmt.Println(res[i])
	}

}

func fileToSlice(asciiFont string) []string {
	file, err := os.Open("justify/asciiTxt/" + asciiFont + ".txt")
	if err != nil {
		log.Fatalf("failed opening file: %s\n Expecting: standard, shadow or thinkertoy", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	return txtlines
}

func addChar(i int, result []string, txtlines []string) []string {
	fmt.Println("length:", len(txtlines[i]))
	for i2 := 0; i2 <= 7; i2++ {
		result[i2] = result[i2] + txtlines[i-1]
		i = i + 1
	}
	return result
}

func AsciiOutput(words, asciiFont, justifyInput string) {
	if !checkAscii(words) {
		log.Fatal("\nCharacter/s in input string is not valid.\n\n" +
			"Please only use characters from the ascii table.\n\n" +
			"Usage: go run . [STRING] [BANNER] [OPTION]\n\n" +
			"Example: go run . something standard --align=right")
	}
	justifyReg := regexp.MustCompile(`^--align=(center|left|right|justify)\z`)
	justifyCheck := justifyReg.MatchString(justifyInput)
	if !justifyCheck {
		log.Fatal("\nInvalid alignment specified.\n\n" +
			"Please only choose one of either 'center', 'left', " +
			"'right' or 'justify'.\n\n" +
			"Usage: go run . [STRING] [BANNER] [OPTION]\n\n" +
			"Example: go run . something standard --align=right")
	}
	//finalLen := getTotalLength(words, asciiFont)
	//fmt.Println(finalLen)
	alignment := justifyInput[8:]
	termLength := getTermLength()
	fmt.Println("terminal length:", termLength)
	spaces := 0
	for _, r := range words {
		if r == ' ' {
			spaces++
		}
	}
	//if spaces > 0 && alignment == "justify" {
	//	termLength = termLength - finalLen
	//	termLength = termLength / spaces
	//}
	fmt.Println("spaces", spaces)
	result := make([]string, 8)
	asciiNbr := 0
	skipN := 0
	txtlines := fileToSlice(asciiFont)
	finalLen := getTotalLength(words, asciiFont, txtlines)
	if finalLen > termLength {
		log.Fatal("Does not fit in terminal.\n\nPlease either shorten your input string or increase the size of the terminal")
	}
	fmt.Println("length of char:", finalLen)
	if spaces > 0 && alignment == "justify" {
		termLength = termLength - finalLen
		termLength = termLength / spaces
	}
	fmt.Println("add to spaces:",termLength)
	for i, r := range words {
		if r == '\\' && words[i+1] == 'n' {
			printAscii(result, alignment, termLength)
			result = make([]string, 8)
			skipN = 2
		} else if skipN == 0 {
			asciiNbr = int(r-' ')*9 + 2
			fmt.Print(string(r) + " ")
			addChar(asciiNbr, result, txtlines)
			if alignment == "justify" {
				if r == ' ' {
					justify(result, termLength)
				}
			}
		}
		if skipN != 0 {
			skipN = skipN - 1
		}
	}
	printAscii(result, alignment, termLength)
}
//you can do it