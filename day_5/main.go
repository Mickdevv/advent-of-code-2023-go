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
	description string 
	table [][3]int
}

func main() {
	file, err := os.Open("day_5/input.txt")
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
			// fmt.Println("Map condition")
			mapTable := MapTable{ description: strings.Split(line, " ")[0] }
			maps = append(maps, mapTable)
		} else if len(line) > 0 {
			// fmt.Println("Line length", line)
			splitLine := strings.Split(line, " ")
			line_numbers_array := [3]int{strToInt(splitLine[0]), strToInt(splitLine[1]), strToInt(splitLine[2])}
			// fmt.Println(len(splitLine))
			// fmt.Println(maps, len(maps))
			maps[len(maps)-1].table = append(maps[len(maps)-1].table, line_numbers_array)
		}
		lineCount += 1
	}
	for _, m := range maps {
		fmt.Println(m)
		fmt.Println()
	}
}

func strToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return i
}