package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type partNumber struct {
	value int
	start int
	end int
	lineNumber int
	toAddRatios bool
}


func main() {

	var slice []string
	var partNumbers []partNumber
	total := 0
	gearRatios := 0

	// LINE_LENGTH := 0

    file, err := os.Open("day_3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	

	scanner := bufio.NewScanner(file)

	
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}
	LINE_LENGTH := len(slice[0])
	LINE_COUNT := len(slice)

	for lineNumber, line := range slice {
		splitLine := strings.Split(removeSymbols(line), ".")
		indexOffset := 0
		for i, v := range splitLine {
			n, err := strconv.Atoi(v)
			if err == nil {
				partNumber := partNumber{value: n, start: i + indexOffset, end: i + indexOffset + len(v), lineNumber: lineNumber, toAddRatios: false}
				indexOffset += len(v)
				partNumbers = append(partNumbers, partNumber)
			}
		}

		for i, v := range line {
			if string(v) == "*" {
				fmt.Println("===========================")
				var adjacentNumbers []int
				// left := i-1
				// if left < 0 {
				// 	left = 0
				// }
				// right := i+1
				// if right > LINE_LENGTH {
				// 	right = LINE_LENGTH
				// }
				// top := lineNumber-1
				// if top < 0 {
				// 	top = 0
				// }
				// bottom := lineNumber +1
				// if bottom > LINE_COUNT {
				// 	bottom = LINE_COUNT
				// }

				if lineNumber > 0 {
					fmt.Println(slice[lineNumber-1])
				}
				fmt.Println(line, i)
				if lineNumber < LINE_COUNT {
					fmt.Println(slice[lineNumber+1])
				}
				for range i {
					fmt.Print(" ")
				}
				fmt.Print("*")
				fmt.Println()

				if lineNumber > 0 {
					adjacentNumbers = append(adjacentNumbers, getNumbersFromLine(slice[lineNumber-1], i)...)
					fmt.Println(adjacentNumbers, len(adjacentNumbers))
				}
				
				adjacentNumbers = append(adjacentNumbers, getNumbersFromLine(line, i)...)
				fmt.Println(adjacentNumbers, len(adjacentNumbers))
				if lineNumber < LINE_COUNT {
					adjacentNumbers = append(adjacentNumbers, getNumbersFromLine(slice[lineNumber+1], i)...)
					fmt.Println(adjacentNumbers, len(adjacentNumbers))
				}


				if len(adjacentNumbers) > 0 {
					if len(adjacentNumbers) > 2 {
						fmt.Println(adjacentNumbers, "==========================================================================================================")
					}
					if len(adjacentNumbers) == 2 {
						fmt.Println(adjacentNumbers, len(adjacentNumbers), adjacentNumbers[0] * adjacentNumbers[1])
						gearRatios += adjacentNumbers[0] * adjacentNumbers[1]
					}

				}
			}
		}
	}

	for _, partNumber := range partNumbers {
		toAdd := false

		endCheck := partNumber.end + 1
		if LINE_LENGTH < partNumber.end + 1 {
			endCheck = LINE_LENGTH
		}

		startCheck := partNumber.start - 1
		if partNumber.start - 1 < 0 {
			startCheck = 0
		}

		if partNumber.lineNumber > 0 {
			if checkForSymbols(slice[partNumber.lineNumber-1][startCheck:endCheck]) {
				toAdd = true
			}
		}		
		if partNumber.lineNumber > 0 {
			if checkForSymbols(slice[partNumber.lineNumber][startCheck:endCheck]) {
				toAdd = true
			}
		}		
		if partNumber.lineNumber < LINE_COUNT -1 {
			if checkForSymbols(slice[partNumber.lineNumber+1][startCheck:endCheck]) {
				toAdd = true
			}
		}

		if toAdd {
			total += partNumber.value
		}
	}

	ln := 2

	for _, part := range partNumbers {
		if part.lineNumber == ln {
			// fmt.Printf("Part Number: %d, Start: %d, End: %d, Line Number: %d\n", part.value, part.start, part.end, part.lineNumber)
		}
	}

	fmt.Println(total)
	fmt.Println(gearRatios)
	// getNumbersFromLine("..........905./50.........@...................971..................762.*..................169........91682..........=......533.......502..", 104)

}

func checkForSymbols(str string) bool{
	for _, c := range str {
		if _, err := strconv.Atoi(string(c)); err != nil && string(c) != "." {
			return true
		}
	}
	return false
}

func isNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func removeSymbols(str string) string {
	newString := ""

	for _, c := range str {
		if !isNumber(string(c)) {
			newString = newString + "."
		} else {
			newString = newString + string(c)
		}
	}

	return newString
}

func getNumbersFromLine(line string, lineIndex int) (result []int) {

	leftNumber := ""
	if lineIndex >= 0 && isNumber(string(line[lineIndex-1])) {
		for i := lineIndex -1; i >= 0 && isNumber(string(line[i])); i-- {
			leftNumber = string(line[i]) + leftNumber
		}
	}
	rightNumber := ""
	if lineIndex <= len(line) && isNumber(string(line[lineIndex+1])) {
		for i := lineIndex + 1; i < len(line) && isNumber(string(line[i])); i++ {
			rightNumber = rightNumber + string(line[i])
		}
	}

	if isNumber(string(line[lineIndex])) {
		finalNumber := leftNumber + string(line[lineIndex]) + rightNumber
		finalNumberInt, _ := strconv.Atoi(finalNumber)

		result = append(result, finalNumberInt)
		fmt.Println("Final", finalNumber)
	} else {
		if leftNumber != "" {
			leftNumberInt, _ := strconv.Atoi(leftNumber)
			result = append(result, leftNumberInt)
		}
		if rightNumber != "" {
			rightNumberInt, _ := strconv.Atoi(rightNumber)
			result = append(result, rightNumberInt)
		}
		
		fmt.Println("Final", leftNumber, rightNumber)
	}
	return 

}