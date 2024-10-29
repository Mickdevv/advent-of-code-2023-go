package main

import (
	"bufio"
	"fmt"
	"os"
)

type cycle struct {
	initial_moves int
	cycle_length int
	cycle_found int
}

func main() {
	file, _ := os.Open("day_8/input.txt")
	scanner := bufio.NewScanner(file)

	maps := make(map[string][2]string)
	var cycles []int

	scanner.Scan()
	moves := scanner.Text()
	scanner.Scan()
	
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			maps[line[:3]] = [2]string{line[7:10], line[12:15]}
		}
	}

	// moveCount_1 := 0
	// currentPos_1 := "AAA"
	// moveCount_2 := 0
	currentPos_2 := []string{}
	for k := range maps {
		if string(k[2]) == "A" {
			currentPos_2 = append(currentPos_2, k)
			cycles = append(cycles, 0)
		}
	}

	fmt.Println(currentPos_2)

	
	
	// for currentPos_1 != "ZZZ" {
	// 	for _, c := range moves {
	// 		if string(c) == "L" {
	// 			currentPos_1 = maps[currentPos_1][0]
	// 		} else if string(c) == "R" {
	// 			currentPos_1 = maps[currentPos_1][1]
	// 		}
	// 		moveCount_1 += 1
	// 	}
	// }

	for i := range currentPos_2 {
		moveCount := 0
		for string(currentPos_2[i][2]) != "Z" {
			move := string(moves[moveCount % len(moves)])
			currentNode := currentPos_2[i]
			currentPos_2[i] = nextNode(maps, currentNode, move)
			moveCount++
		}
		cycles[i] = moveCount

	}
	fmt.Println(cycles)
	fmt.Println(lcmOfSlice(cycles))
		

}

func checkArrival(positions []string) bool {
	arrived := true
	for _, pos := range positions {
		if string(pos[2]) != "Z" {
			arrived = false
		}
	}
	return arrived
}

// Helper function to calculate the Greatest Common Divisor (GCD) of two numbers
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

// Function to calculate the Least Common Multiple (LCM) of two numbers
func lcm(a, b int) int {
    return a * (b / gcd(a, b))
}

// Function to calculate the LCM of a slice of integers
func lcmOfSlice(numbers []int) int {
    if len(numbers) == 0 {
        return 0 // Return 0 if the slice is empty
    }

    result := numbers[0]
    for _, num := range numbers[1:] {
        result = lcm(result, num)
    }
    return result
}

func nextNode(maps map[string][2]string, currentNode string, move string) string {
	if move == "L" {
		return maps[currentNode][0]
	} else if move == "R" {
		return maps[currentNode][1]
	}
	return ""
}