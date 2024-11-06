package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("day_11/input.txt")
	scanner := bufio.NewScanner(file)


	var universe []string
	lineIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		if !containsGalaxies(line) {
			newLine := ""
			for range len(line) {
				newLine += "*"
			}
			universe = append(universe, newLine)
			} else {
			
			universe = append(universe, line)
		}
		lineIndex ++
	}
	// showUniverse(expandUniverse(universe))

	universe = expandUniverse(universe)

	P1(universe)
	P2(universe)

}

func P1(universe []string) {

	galaxies := getGalaxyPositions(universe, 2)
	
	distance_sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i; j < len(galaxies); j++ {
			distance_sum += getGalaxyDistance(galaxies[i], galaxies[j])
		}
	}
	
	fmt.Println("galaxies : ", len(galaxies), " | Possible combinations : ", len(galaxies) * (len(galaxies)-1) / 2)
	fmt.Println(distance_sum)
}

func P2(universe []string) {
		var galaxies [][2]int


	galaxies = getGalaxyPositions(universe, 1000000)
	
	distance_sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i; j < len(galaxies); j++ {
			distance_sum += getGalaxyDistance(galaxies[i], galaxies[j])
		}
	}
	
	fmt.Println("galaxies : ", len(galaxies), " | Possible combinations : ", len(galaxies) * (len(galaxies)-1) / 2)
	fmt.Println(distance_sum)
}

func containsGalaxies(line string) bool {
	for _, c := range line {
		if string(c) == "#" {
			return true
		}
	}
	return false
} 

func showUniverse(universe []string) {
	for _, line := range universe {
		fmt.Println(line)
	}
}

func expandUniverse(universe []string) []string {
	
	for charIndex := range universe[0] {
		galaxyFound := false
		for _, line := range universe {
			if string(line[charIndex]) == "#" {
				galaxyFound = true
				break
			}
		}
		if !galaxyFound {
			for index, line := range universe {
				newLine := []rune(line)
				newLine[charIndex] = '*'
				// fmt.Println(string(newLine))
				universe[index] = string(newLine)
			}
		}
	}
	return universe
}

func getGalaxyPositions(universe []string, expansionCoefficient int) [][2]int {
	expansionCoefficient = expansionCoefficient -1
	var galaxies [][2]int
	currentPositionOffset := [2]int{0,0}
	for y, line := range universe {
		currentPositionOffset[0] = 0
		if isLineExpanded(line) {
			currentPositionOffset[1] += expansionCoefficient
		} else {
			for x, c := range line {
				if string(c) == "#" {
					galaxies = append(galaxies, [2]int{x + currentPositionOffset[0], y + currentPositionOffset[1]})
					// fmt.Println([2]int{x + currentPositionOffset[0], y + currentPositionOffset[1]})
				} else if string(c) == "*" {
					currentPositionOffset[0] += expansionCoefficient
				}
			}

		}
	}
	// fmt.Println(galaxies)
	return galaxies
}

func getGalaxyDistance(galaxy1 [2]int, galaxy2 [2]int) int {
	return abs(galaxy1[0] - galaxy2[0]) + abs(galaxy1[1] - galaxy2[1])
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func getEmptyLines(universe []string) [2][]int {
	var emptyLines [2][]int

	charIndexOffset := 0
	for charIndex := range universe[0] {
		galaxyFound := false
		for _, line := range universe {
			if string(line[charIndex + charIndexOffset]) == "#" {
				galaxyFound = true
				break
			}
		}
		if !galaxyFound {
			for index, line := range universe {
				universe[index] = line[:charIndex + charIndexOffset] + "." + line[charIndex + charIndexOffset:]
			}
			charIndexOffset ++
		}
	}
	return emptyLines
}

func isLineExpanded(line string) bool {
	for _, c := range line {
		if string(c) != "*" {
			return false
		}
	}
	return true
}