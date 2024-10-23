package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {

	numbersMap := map[string]string{"zero":"0", "one":"1", "two":"2", "three":"3", "four":"4", "five":"5", "six":"6", "seven":"7", "eight":"8", "nine":"9"}

    file, err := os.Open("day_1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		
		line := scanner.Text()
		line_slice := []string{}

		for i, c := range line {
			if unicode.IsDigit(c) {
				line_slice = append(line_slice, string(c))
			} else {
				for key, val := range numbersMap {
					inc := i+len(key)
					if inc > len(line) {
						inc = len(line)
					}
					if line[i:inc] == key {
						line_slice = append(line_slice, val)
					}
				}
			}
		}

		digits, err := strconv.Atoi(line_slice[0] + line_slice[len(line_slice)-1])
		if err != nil {
			log.Fatal(err)
		}

		total += digits
		fmt.Printf("%s, %s, %d, %d\n", line_slice, line, digits, total)
	}
	fmt.Println(total)
	
}

func extractDigits(textLine string) int {
	firstDigit, lastDigit := "", ""
	for _, c := range textLine {
		if unicode.IsDigit(c) {
			if firstDigit == "" {
				firstDigit = string(c)
			}
			lastDigit = string(c)
		}
	}

	result, err := strconv.Atoi(firstDigit + lastDigit)
	if err != nil {
		log.Fatal(err)
	}
	return result
}