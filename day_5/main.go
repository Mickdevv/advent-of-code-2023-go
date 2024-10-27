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
	var maps []MapTable
	lineCount := 0
	
	scanner.Scan()
	line := scanner.Text()
	// fmt.Println(line)
	
	for _, v := range strings.Split(strings.Split(line, ": ")[1], " ") {
		v_int, _ := strconv.Atoi(v)
		seeds = append(seeds, v_int)
	}
	
	// fmt.Println(seeds)
	
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

	seedIterations := 0
	mapIterations := 0
	rowIterations := 0
	currentNumberCheckIterations := 0
	currentNumberCheckFound := 0

	
	lowest_location_number := -1
	// var seeds_part2 []int
	for i := 0; i < len(seeds); i+=2 {
		fmt.Println("New seed range : ", (i+2)/2, "/", len(seeds)/2, " | ", seeds[i], seeds[i+1])
		for j := seeds[i]; j <= seeds[i] + seeds[i+1]; j++ {
			// seedIterations ++
			currentNumber := j
			for _, currentMap := range maps {
				// mapIterations ++
				for _, row := range currentMap.table {
					// rowIterations ++
					// rowEnd := row[1] + row[2]
					if currentNumber >= row[1] {
						// currentNumberCheckIterations ++
						if currentNumber <= row[3] {
							// currentNumberCheckFound ++
							currentNumber = row[0] + (currentNumber - row[1])
							break
						}
					}
				}
			}
			if currentNumber < lowest_location_number || lowest_location_number == -1 {
				lowest_location_number = currentNumber
				fmt.Println("New lowest number found : ", currentNumber)
			}
		// }
		}
	}


	fmt.Println("Seed iterations : ", seedIterations, " | Map iterations : ", mapIterations, " | Row iterations : ", rowIterations, " | currentNumber check 1 : ", currentNumberCheckIterations, " | currentNumber check 2 : ", currentNumberCheckFound)

	fmt.Println(lowest_location_number)
}

// func processValueTable(table [][3]int, input int) int {
	
// }

func strToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func findMapBySource(source string, mapArray []MapTable) (MapTable) {
	for _, m := range mapArray {
		if m.source == source {
			return m
		}
	}
	// fmt.Println("Map not found : ", source)
	return MapTable{}
}