package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("day_11/input.txt")
	scanner := bufio.NewScanner(file)

	var stars [][2]int

	line_number := 0

	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			if string(c) == "#" {
				stars = append(stars, [2]int{i, line_number})
			}
		}
		line_number ++
	}

	fmt.Println("Stars : ", len(stars), " | Possible combinations : ", len(stars) * (len(stars)-1) / 2)
	
}