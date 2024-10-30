package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var histories [][]int
	file, _ := os.Open("day_9/input_test.txt")

	scanner := bufio.NewScanner(file)


	for scanner.Scan() {
		line := scanner.Text()
		var history []int
		for _, n := range strings.Split(line, " ") {
			to_add, _ :=  strconv.Atoi(n)
			history = append(history, to_add)
		}
		histories = append(histories, history)

	}

	fmt.Println(histories)

}
