package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type MapTable struct {
	source string 
	destination string 
	table [][4]int
}


func main() {
	file, err := os.Open("day_5/input.txt")
	// file, err := os.Open("day_5/input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	
	var seeds []int
	var seedsP2 [][2]int
	var maps []MapTable
	lineCount := 0
	
	scanner.Scan()
	line := scanner.Text()
	for _, v := range strings.Split(strings.Split(line, ": ")[1], " ") {
		v_int, _ := strconv.Atoi(v)
		seeds = append(seeds, v_int)
	}
	
	
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "map") {
			mapTable := MapTable{ source: strings.Split(strings.Split(line, " ")[0], "-")[0], destination: strings.Split(strings.Split(line, " ")[0], "-")[2] }
			maps = append(maps, mapTable)
		} else if len(line) > 0 {
			splitLine := strings.Split(line, " ")
			endRange := strToInt(splitLine[1]) + strToInt(splitLine[2])
			line_numbers_array := [4]int{strToInt(splitLine[0]), strToInt(splitLine[1]), strToInt(splitLine[2]), endRange}
			maps[len(maps)-1].table = append(maps[len(maps)-1].table, line_numbers_array)
		}
		lineCount += 1
	}

	
	// lowest_location_number := -1
	for i := 0; i < len(seeds); i+=2 {
		seedsP2 = append(seedsP2, [2]int{seeds[i], seeds[i] + seeds[i+1]})
	}

	// fmt.Println(seedsP2)

	for _, currentMap := range maps {
		fmt.Println(currentMap.source, currentMap.destination)

		var new [][2]int
		for len(seedsP2) > 0 {
			seed := seedsP2[0]
			// fmt.Println(len(seedsP2))
			if len(seedsP2) > 0 {
				seedsP2 = seedsP2[1:]
			}
			matched := false

			for _, row := range currentMap.table {
				overlapStart := max(seed[0], row[1])
				overlapEnd := min(seed[1], row[1] + row[2])
				if overlapStart < overlapEnd {
					new = append(new, [2]int{overlapStart - row[1] + row[0], overlapEnd - row[1] + row[0]})
					if overlapStart > seed[0] {
						seedsP2 = append(seedsP2, [2]int{ seed[0], overlapStart})
					}
					if overlapEnd < seed[1] {
						seedsP2 = append(seedsP2, [2]int{ overlapEnd, seed[1]})
					}
					matched = true
					break
				}
			}
			if !matched {
				new = append(new, seed)
			}
		}
		seedsP2 = new
		
	}
	lowest_location_number := seedsP2[0][0]
	for _, seed := range seedsP2 {
		if seed[0] < lowest_location_number {
			lowest_location_number = seed[0]
		}
	}
	fmt.Println(lowest_location_number)
	// for i := 0; i < len(seeds); i+=2 {
	// 	fmt.Println("New seed range : ", (i+2)/2, "/", len(seeds)/2, " | ", seeds[i], seeds[i+1])
	// 	for j := seeds[i]; j <= seeds[i] + seeds[i+1]; j++ {
	// 		currentNumber := applyMaps(j, maps)
			
	// 		if currentNumber < lowest_location_number || lowest_location_number == -1 {
	// 			lowest_location_number = currentNumber
	// 			fmt.Println("New lowest number found : ", currentNumber)
	// 		}
	// 	// }
	// 	}
	// }


	// fmt.Println(lowest_location_number)
}

func strToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func applyMaps(seed int, maps []MapTable) int {
	currentNumber := seed
	for _, currentMap := range maps {
		for _, row := range currentMap.table {
			if currentNumber >= row[1] {
				if currentNumber <= row[3] {
					currentNumber = row[0] + (currentNumber - row[1])
					break
				}
			}
		}
	}
	return currentNumber

}