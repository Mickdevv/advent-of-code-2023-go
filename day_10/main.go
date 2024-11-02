package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

var move_map = map[string][2][2]int{
	"|": {{0, 1}, {0, -1}}, // up/down
	"-": {{-1, 0}, {1, 0}}, // left/right

	"L": {{1, 0}, {0, -1}}, // up-left/down-right
	"J": {{-1, 0}, {0, -1}}, // down-left/up-right
	"7": {{-1, 0}, {0, 1}}, //down-right/up-left
	"F": {{1, 0}, {0, 1}}, //down-left/up-right

	"S": {{0, 0}, {0, 0}}, //down-left/up-right
}

func main() {

	var grid []string
	// file, _ := os.Open("day_10/input_test.txt")
	file, _ := os.Open("day_10/input.txt")
	scanner := bufio.NewScanner(file)

	startingPoint := [2]int{0,0}
	var startingPointX int
	startingPointY := 0

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
		startingPointX = sliceContains(line, "S")
		if startingPointX != -1 {
			startingPoint = [2]int{startingPointX, startingPointY}
			// fmt.Println(startingPoint)
		}
		startingPointY ++
	}

	shapePoints := P1(startingPoint, grid)
	
	P2(startingPoint, grid, shapePoints)
}

func P2(startingPoint [2]int, grid []string, shapePoints [][2]int) {
	count := 0
	for i, line := range grid {
		lineAbove := ""
		lineBelow := ""
		if i > 0 {
			lineAbove = grid[i-1]
		}
		if i < len(grid)-2 {
			lineBelow = grid[i+1]
		}
		// fmt.Println(lineAbove)
		// fmt.Println(line)
		// fmt.Println(lineBelow)
		fmt.Println("----", i)
		fmt.Println()
		count += rayCastLine(lineAbove, line, lineBelow, shapePoints, i)
	}

	fmt.Println(count)
}

func P1(startingPoint [2]int, grid []string) [][2]int {
	var shapePoints [][2]int
	moveCounter := 1
	previousPos, currentPos := findFirstMove(startingPoint, grid, startingPoint)
	shapePoints = append(shapePoints, previousPos)
	shapePoints = append(shapePoints, currentPos)
	
	for string(grid[currentPos[1]][currentPos[0]]) != "S" {
		previousPos, currentPos = findNextMove(currentPos, grid, previousPos)
		shapePoints = append(shapePoints, currentPos)
		moveCounter ++ 
	}
	fmt.Println(math.Ceil(float64(moveCounter) / 2))

	return shapePoints

}

func sliceContains(str string, char string) int {
	for i, c := range str {
		if string(c) == char {
			return i
		}
	}
	return -1
}

func findFirstMove(currentPos [2]int, grid []string, previousPos [2]int) ([2]int, [2]int) {
	for y := max(currentPos[1]-1, 0); y < min(currentPos[1]+2, len(grid)); y++ { //Y
		for x := max(currentPos[0]-1, 0); x < min(currentPos[0]+2, len(grid[currentPos[1]])); x++ { //X
			// fmt.Println(string(grid[y][x]), y, x)
			if nextVal, ok := move_map[string(grid[y][x])]; ok {
				// fmt.Println(string(grid[y][x]), val, isMoveValid(startingPoint, [2]int{x, y}, val), [2]int{x, y}, previousPos)
				nextPos := [2]int{x, y}
				currentVal := move_map[string(grid[currentPos[1]][currentPos[0]])]
				if (isMoveValid(currentPos, nextPos, nextVal, currentVal) && nextPos != previousPos) {
					// fmt.Println(previousPos, string(grid[previousPos[1]][previousPos[0]]), currentPos, string(grid[currentPos[1]][currentPos[0]]), nextPos, string(grid[nextPos[1]][nextPos[0]]))
					return currentPos, nextPos
				}
			}
		}
	}
	fmt.Println("----FAIL----")
	fmt.Println(previousPos, string(grid[previousPos[1]][previousPos[0]]), currentPos, string(grid[currentPos[1]][currentPos[0]]))
	log.Fatal("No valid next move found")
	return [2]int{-1,-1}, [2]int{-1,-1}
}

func findNextMove(currentPos [2]int, grid []string, previousPos [2]int) ([2]int, [2]int) {
	currentChar := string(grid[currentPos[1]][currentPos[0]])
	currentCharOffsets := move_map[currentChar]

	m1 := [2]int{currentPos[0] + currentCharOffsets[0][0], currentPos[1] + currentCharOffsets[0][1]}
	m2 := [2]int{currentPos[0] + currentCharOffsets[1][0], currentPos[1] + currentCharOffsets[1][1]}

	// fmt.Println(previousPos, currentPos, m1, m2, currentChar)

	if m1 != previousPos {
		return currentPos, m1
	} else if m2 != previousPos {
		return currentPos, m2
	}
	log.Fatal("No valid next move found")
	return [2]int{-1,-1}, [2]int{-1,-1}
	
}

func isMoveValid(currentPos [2]int, nextCharPos [2]int, nextVal [2][2]int, currentVal [2][2]int) bool {

	m1 := addCoordinates(nextCharPos, nextVal[0])
	m2 := addCoordinates(nextCharPos, nextVal[1])

	if currentPos == m1 || currentPos == m2 {
		return true
	}
	return false

}

func addCoordinates(pos1 [2]int, pos2 [2]int) [2]int {
	return [2]int{pos1[0] + pos2[0], pos1[1] + pos2[1]}
}

// ray casting algorithm
func rayCastLine(lineAbove string, line string, lineBelow string, shapePoints [][2]int, lineIndex int) (int) {
	insideBoundary := false
	positionsFound := 0
	partOfShape := false

	specialCharSeq := false
	sChars := [2]string{"", ""}

	for i, c := range line {
		char := string(c)
		partOfShape = false

		if isSpecialChar(char) || char == "|" {
			partOfShape = isPartOfShape(shapePoints, [2]int{i, lineIndex})
		}

		if partOfShape {

			if char == "S" {
				if len(lineBelow) > 0 && len(lineAbove) > 0 && isUpMove(string(lineAbove[i])) && isDownMove(string(lineBelow[i])) {
					insideBoundary = !insideBoundary
					// fmt.Println(sChars)
				} else if len(lineAbove) > 0 && isUpMove(string(lineAbove[i])) && specialCharSeq == false {
					specialCharSeq = true
					sChars[0] = "L"
					// fmt.Println(sChars)
				} else if len(lineBelow) > 0 && isDownMove(string(lineBelow[i])) && specialCharSeq == false {
					specialCharSeq = true
					sChars[0] = "F"
					// fmt.Println(sChars)
				} else if len(lineAbove) > 0 && isUpMove(string(lineAbove[i])) && specialCharSeq == true {
					specialCharSeq = false
					sChars[1] = "J"
					// fmt.Println(sChars)
				} else if len(lineBelow) > 0 && isDownMove(string(lineBelow[i])) && specialCharSeq == true {
					specialCharSeq = false
					sChars[1] = "7"
					// fmt.Println(sChars)
				}
			} else if char == "|" {
				insideBoundary = !insideBoundary
			} else if isSpecialChar(char) && char != "S" {
				if char == "L" || char == "F" {
					specialCharSeq = true
					sChars[0] = char
					// fmt.Println(sChars)
				} else if char == "J" || char == "7" {
					specialCharSeq = false
					sChars[1] = char
					// fmt.Println(sChars)
				}
			}
			if sChars[0] != "" && sChars[1] != "" {
				// fmt.Println(line)
				// fmt.Println(sChars)
				if isCharSequenceBoundary(sChars) {
					insideBoundary = !insideBoundary
				}
				fmt.Println(sChars, isCharSequenceBoundary(sChars), insideBoundary)
				sChars = [2]string{"", ""}
			}
		}
		
		if insideBoundary && !partOfShape {
			fmt.Println(lineAbove)
			fmt.Println(line, char, insideBoundary)
			fmt.Println(lineBelow)
			fmt.Println()
			positionsFound ++
		}
	}
	return positionsFound
}

func isSpecialChar(char string) bool {
	if char == "S" ||char == "-" ||char == "L" ||char == "J" ||char == "F" ||char == "7" {
		return true
	} else {
		return false
	}
}

func isCharSequenceBoundary(sChars [2]string) bool {
	// fmt.Println(sChars)

	if sChars == [2]string{"L", "J"} || sChars == [2]string{"F", "7"} {
		return false
	} else if sChars == [2]string{"F", "J"} || sChars == [2]string{"L", "7"} {
		return true
	} else {
		fmt.Println("Something went wrong", sChars)
		return false
	}
}

func isPartOfShape(shapePoints [][2]int, point [2]int) bool {
	for _, v := range shapePoints {
		if v == point {
			return true
		}
	}
	return false

}

func findCharInLine(line string, char string) int {
	for i, c := range line {
		if char == string(c) {
			return i
		}
	}
	return -1
}

func isUpMove(c string) bool {
	if c == "|" || c == "S" || c == "F" || c == "7" {
		return true
	}
	return false
}


func isDownMove(c string) bool {
	if c == "|" || c == "S" || c == "L" || c == "J" {
		return true
	}
	return false
}
