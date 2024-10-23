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
}


func main() {
	var slice []string
	var partNumbers []partNumber
	total := 0

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
				partNumber := partNumber{value: n, start: i + indexOffset, end: i + indexOffset + len(v), lineNumber: lineNumber}
				indexOffset += len(v)
				partNumbers = append(partNumbers, partNumber)
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
			fmt.Printf("Part Number: %d, Start: %d, End: %d, Line Number: %d\n", part.value, part.start, part.end, part.lineNumber)
		}
	}

	fmt.Println(slice[ln])
	
	fmt.Println(total)

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