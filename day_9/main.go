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
	file, _ := os.Open("day_9/input.txt")
	// file, _ := os.Open("day_9/input_test.txt")

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

	P2(histories)


}

func P1(histories [][]int) {
		
	acc := 0
	for _, v := range histories {
		acc += findNextValue(v)
	}
	fmt.Println(acc)

}

func P2(histories [][]int) {
	acc := 0
	for _, v := range histories {
		acc += findPreviousValue(v)
	}
	fmt.Println(acc)
}

func findNextValue(history []int) int {
	history_table := [][]int{ history }
	lastLine := history

	for !checkAllZeroes(lastLine) {
		var newLine []int
		for i := range lastLine {
			if i < len(lastLine) -1 {
				newValue := lastLine[i+1] - lastLine[i]
				newLine = append(newLine, newValue)
			}
		}
		history_table = append(history_table, newLine)
		lastLine = history_table[len(history_table)-1]
	}
	
	history_table[len(history_table)-1] = append(history_table[len(history_table)-1], 0)

	// starts at 2nd to last element and walks backwards from there
	for i := len(history_table)-2; i >= 0; i -- {
		lastElementCurrentLine := history_table[i][len(history_table[i])-1] // 2
		lastElementPreviousLine := history_table[i+1][len(history_table[i+1])-1] // 0

		history_table[i] = append(history_table[i], lastElementCurrentLine + lastElementPreviousLine )

		
	}
	// fmt.Println(history_table)

	return history_table[0][len(history_table[0])-1]
	
}

func findPreviousValue(history []int) int {
	history_table := [][]int{ history }
	lastLine := history

	for !checkAllZeroes(lastLine) {
		var newLine []int
		for i := range lastLine {
			if i < len(lastLine) -1 {
				newValue := lastLine[i+1] - lastLine[i]
				newLine = append(newLine, newValue)
			}
		}
		history_table = append(history_table, newLine)
		lastLine = history_table[len(history_table)-1]
	}
	
	history_table[len(history_table)-1] = append(history_table[len(history_table)-1], 0)

	// starts at 2nd to last element and walks backwards from there
	for i := len(history_table)-2; i >= 0; i -- {
		firstElementCurrentLine := history_table[i][0] // 2
		firstElementPreviousLine := history_table[i+1][0] // 0

		history_table[i] = append([]int{firstElementCurrentLine - firstElementPreviousLine}, history_table[i]...)

		
	}
	// fmt.Println(history_table)

	return history_table[0][0]
	
}

func checkAllZeroes(history []int) bool {
	for _, v := range history {
		if v != 0 {
			return false
		}
	}
	return true
}