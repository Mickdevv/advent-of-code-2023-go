package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("day_11/input.txt")
	scanner := bufio.NewScanner(file)

	var galaxies [][2]int
	var universe []string

	for scanner.Scan() {
		line := scanner.Text()
		universe = append(universe, line)
		if !containsGalaxies(line) {
			universe = append(universe, line)
		}
	}
	// showUniverse(universe)
	universe = expandUniverse(universe)
	
	// showUniverse(universe)
	// fmt.Println(len(universe[0]))

	galaxies = getGalaxyPositions(universe)
	
	distance_sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i; j < len(galaxies); j++ {
			// fmt.Println(galaxies[i], galaxies[j])
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
	return universe
}

func getGalaxyPositions(universe []string) [][2]int {
	var galaxies [][2]int
	for y, line := range universe {
		for x, c := range line {
			if string(c) == "#" {
				galaxies = append(galaxies, [2]int{x, y})
			}
		}
	}
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