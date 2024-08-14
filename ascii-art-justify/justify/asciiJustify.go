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

func countSpace(words string) (int, bool) {
	spaces := 0
	for in, ru := range words {
		if ru == '\\' && words[in+1] == 'n' {
			break
		}
		if ru == ' ' {
			spaces++
		}
	}
	if spaces > 0 {
		return spaces, false
	} else {
		for i2, r2 := range words {
			if r2 == '\\' && words[i2+1] == 'n' {
				break
			} else {
				spaces++
			}
		}
		spaces = spaces - 1
		return spaces, true
	}
}

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
	for i, r := range words {
		if !(r == '\\' && words[i+1] == 'n') {
			lenOfAscii = len(txtlines[int(r-' ')*9+2])
			result = result + lenOfAscii
		} else {
			return result
		}
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
	for i2 := 0; i2 <= 7; i2++ {
		result[i2] = result[i2] + txtlines[i-1]
		i = i + 1
	}
	return result
}

func AsciiJustify(words, asciiFont, justifyInput string) {
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
	alignment := justifyInput[8:]
	termLength := getTermLength()
	spaces := 0
	noSpace := false
	spaces, noSpace = countSpace(words)
	result := make([]string, 8)
	asciiNbr := 0
	skipN := 0
	txtlines := fileToSlice(asciiFont)
	finalLen := getTotalLength(words, asciiFont, txtlines)
	if finalLen > termLength {
		log.Fatal("Does not fit in terminal.\n\nPlease either shorten your input string or increase the size of the terminal")
	}
	if spaces > 0 && alignment == "justify" {
		termLength = termLength - finalLen
		termLength = termLength / spaces
	}
	for i, r := range words {
		if r == '\\' && words[i+1] == 'n' {
			if alignment == "justify" {
				noSpace = false
				printAscii(result, alignment, termLength)
				termLength = getTermLength()
				termLength = termLength - getTotalLength(words[i+2:], asciiFont, txtlines)
				spaces, noSpace = countSpace(words[i+2:])
				termLength = termLength / spaces
				fmt.Println()
			} else {
				printAscii(result, alignment, termLength)
			}
			result = make([]string, 8)
			skipN = 2
		} else if skipN == 0 {
			asciiNbr = int(r-' ')*9 + 2
			//fmt.Print(string(r) + " ")
			addChar(asciiNbr, result, txtlines)
			if alignment == "justify" {
				if (r == ' ' || noSpace) && spaces > 0 {
					justify(result, termLength)
					spaces--
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
