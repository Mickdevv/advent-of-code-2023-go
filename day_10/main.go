package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var move_map = map[string][2][2]int{
	"|": {{0, 1}, {0, -1}}, // up/down
	"-": {{1, 0}, {-1, 0}}, // left/right
	"L": {{-1, 0}, {0, 1}}, // up-left/down-right
	"J": {{1, 0}, {0, 1}}, // down-left/up-right
	"7": {{1, 0}, {0, -1}}, //down-right/up-left
	"F": {{-1, 0}, {0, -1}}, //down-left/up-right
	"S": {{0, 0}, {0, 0}}, //down-left/up-right
}

func main() {

	var grid []string
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
			fmt.Println(startingPoint)
		}
		startingPointY ++
	}

	P1(startingPoint, grid)



}

func P1(startingPoint [2]int, grid []string) {
	moveCounter := 1
	currentPos, previousPos := findNextMove(startingPoint, grid, startingPoint)
	// fmt.Println(currentPos, previousPos)
	for string(grid[currentPos[1]][currentPos[0]]) != "S" {
		currentPos, previousPos = findNextMove(currentPos, grid, previousPos)
		moveCounter ++ 
		fmt.Println(moveCounter)
	}
	fmt.Println(moveCounter)
}

func sliceContains(str string, char string) int {
	for i, c := range str {
		if string(c) == char {
			return i
		}
	}
	return -1
}

func findNextMove(currentPos [2]int, grid []string, previousPos [2]int) ([2]int, [2]int) {
	for y := max(currentPos[1]-1, 0); y < min(currentPos[1]+2, len(grid)); y++ { //Y
		for x := max(currentPos[0]-1, 0); x < min(currentPos[0]+2, len(grid[currentPos[1]])); x++ { //X
			// fmt.Println(string(grid[y][x]), y, x)
			if nextVal, ok := move_map[string(grid[y][x])]; ok {
				// fmt.Println(string(grid[y][x]), val, isMoveValid(startingPoint, [2]int{x, y}, val), [2]int{x, y}, previousPos)
				nextPos := [2]int{x, y}
				currentVal := move_map[string(grid[currentPos[1]][currentPos[0]])]
				if (isMoveValid(currentPos, nextPos, nextVal, currentVal) && nextPos != previousPos) {
					fmt.Println(previousPos, string(grid[previousPos[1]][previousPos[0]]), currentPos, string(grid[currentPos[1]][currentPos[0]]), nextPos, string(grid[nextPos[1]][nextPos[0]]))
					return nextPos, currentPos
				}
			}
		}
	}
	fmt.Println("----FAIL----")
	fmt.Println(previousPos, string(grid[previousPos[1]][previousPos[0]]), currentPos, string(grid[currentPos[1]][currentPos[0]]))
	log.Fatal("No valid next move found")
	return [2]int{-1,-1}, [2]int{-1,-1}
}

func isMoveValid(currentPos [2]int, nextCharPos [2]int, nextVal [2][2]int, currentVal [2][2]int) bool {
	nextCharOffset := [2]int{nextCharPos[0] - currentPos[0], nextCharPos[1] - currentPos[1]}
	currentCharOffset := [2]int{currentPos[0] - nextCharPos[0], currentPos[1]  - nextCharPos[1]}

	if (nextCharOffset == nextVal[0] || nextCharOffset == nextVal[1]) && (currentCharOffset == currentVal[0] || currentCharOffset == currentVal[1] || currentVal[0] == [2]int{0,0}){
		fmt.Println(nextCharOffset, nextVal, currentCharOffset, currentVal, currentPos)
		return true
	}
	return false

}