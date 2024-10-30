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
	file, _ := os.Open("day_10/input_test.txt")
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

	P1(startingPoint, grid)
	
	P2(startingPoint, grid)


}

func P2(startingPoint [2]int, grid []string) {
	count := 0
	for i, line := range grid {
		count += rayCastLine(line)
		fmt.Println("----", i)
	}

	fmt.Println(count)
}

func P1(startingPoint [2]int, grid []string) {
	moveCounter := 1
	previousPos, currentPos := findFirstMove(startingPoint, grid, startingPoint)
	// fmt.Println(previousPos, currentPos)
	
	for string(grid[currentPos[1]][currentPos[0]]) != "S" {
		previousPos, currentPos = findNextMove(currentPos, grid, previousPos)
		moveCounter ++ 
	}
	fmt.Println(math.Ceil(float64(moveCounter) / 2))

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
func rayCastLine(line string) (int) {
	sideCount := 0
	positionsFound := 0

	specialCharSeq := [2]bool{false, false}

	for _, c := range line {

		specialCharSeq[1] = specialCharSeq[0]
		if isSpecialChar(string(c)) {
			specialCharSeq[0] = true
		} else {
			specialCharSeq[0] = false
		}
		if string(c) == "|" {
			fmt.Println(string(c), 1)
			sideCount ++
		} else if (specialCharSeq[0] == true && specialCharSeq[1] == false) && sideCount % 2 == 1 {
			fmt.Println(string(c), 2)
			sideCount ++
		}
		if sideCount > 0 && sideCount % 2 == 1 && !isSpecialChar(string(c)) && string(c) != "|" {
			fmt.Println(string(c), sideCount)
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