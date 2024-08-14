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
	nbr, _ := strconv.Atoi(str[ind:])
	return nbr
}

func getNbrLines(words string) int {
	nbrLines := 1
	for i, r := range words {
		if r == '\\' && words[i+1] == 'n' {
			nbrLines = nbrLines + 1
		}
	}
	return nbrLines
}

func getCharEachLine(words string, nbrLines int) []int {
	charEachLine := make([]int, nbrLines)
	line := 0
	for i := 0; i < len(words); i++ {
		// Skip new line characters in character count
		if words[i] == '\\' && words[i+1] == 'n' {
			line++
			i++
		} else {
			charEachLine[line] = charEachLine[line] + 1
		}
	}
	return charEachLine
}

func getWordsPerLine(words string, nbrLines int) []int {
	wordsPerLine := make([]int, nbrLines)
	for i := 0; i < nbrLines; i++ {
		wordsPerLine[i] = 1
		for j := 0; j < len(words); j++ {
			if words[j] == '\\' && words[j+1] == 'n' {
				i++
				j++
			} else if words[j] == ' ' && words[j+1] != ' ' {
				wordsPerLine[i] = wordsPerLine[i] + 1
			}
		}
	}
	return wordsPerLine
}

func getLineLength(words string, charEachLine []int, nbrLines int) []int {
	lineLength := make([]int, nbrLines)
	line := 0
	for i := 0; i < len(words); i++ {
		if words[i] == '\\' && words[i+1] == 'n' {
			line = line + 1
		} else {
			lineLength[line] = lineLength[line] + 1
		}
	}
	return lineLength
}

func checkAscii(str string) bool {
	for _, r := range str {
		if !(r >= ' ' && r <= '~') {
			return false
		}
	}
	return true
}

func printAscii(result [][]string, charEachLine []int, nbrLines int) {
	for i := 0; i < nbrLines; i++ {
		if result[i][0] == "" {
			fmt.Println()
			return
		}
		for j := 0; j < 8; j++ {
			for k := 0; k < charEachLine[i]; k++ {
				fmt.Print(result[j][k])
			}
			fmt.Println()
		}
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

func addChar(asciiNbr, characterIndex, lineNbr int, result [][]string, txtlines []string, alignment string) [][]string {
	for i := 0; i <= 7; i++ {
		result[i][characterIndex] = result[i][characterIndex] + txtlines[asciiNbr-1]
		asciiNbr = asciiNbr + 1
	}
	return result
}

func addJustification(result [][]string, characterTotal, termLength, nbrLines int, charEachLine, wordsPerLine, lineLength []int, alignment string) [][]string {
	var padding int
	charCount := 0
	switch alignment {
	case "left":
		return result
	case "right":
		for i := 0; i < nbrLines; i++ {
			padding = (termLength - lineLength[i])
			for j := 0; j < charEachLine[i]; j++ {
				for k := 0; k < 8; k++ {
					result[charCount][k] = result[charCount][k] + " "
				}
			}
			charCount++
		}
	case "center":
		for i := 0; i < nbrLines; i++ {
			padding = (termLength - charEachLine[i]) / 2
			for j := 0; j < padding; j++ {
				result[0][j] = " "
			}
		}
	case "justify":
		for i := 0; i < nbrLines; i++ {
			if wordsPerLine[i] == 1 {
				padding = termLength - charEachLine[i]
				for j := 0; j < padding; j++ {
					result[0][j] = " "
				}
			} else {
				padding = (termLength - charEachLine[i]) / (wordsPerLine[i] - 1)
				for j := 0; j < padding; j++ {
					result[0][j] = " "
				}
			}
	}
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
	txtlines := fileToSlice(asciiFont)
	alignment := justifyInput[8:]
	termLength := getTermLength()
	nbrLines := getNbrLines(words)
	charEachLine := getCharEachLine(words, nbrLines)
	characterTotal := 0
	for _, value := range charEachLine {
		characterTotal += value
	}
	// Create a slice of slices to store the ascii characters
	result := make([][]string, characterTotal)
	for i := 0; i < characterTotal; i++ {
		result[i] = make([]string, 8)
	}
	var asciiNbr int
	characterIndex := 0
	lineNbr := 0
	for i := 0; i < len(words); i++ {
		if words[i] == '\\' && words[i+1] == 'n' {
			i++ // Skip the 'n' in '\n'
			lineNbr++
		} else {
			asciiNbr = int(words[i]-' ')*9 + 2 // The first 2 lines of the ascii font are not used
			result = addChar(asciiNbr, characterIndex, lineNbr, result, txtlines, alignment)
			characterIndex++
		}
	}
	printAscii(result, charEachLine, nbrLines)
}
