package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)


func main() {
	mult := 1
	var times []float64
	var distances []float64
	file, err := os.Open("day_6/input.txt")
	// file, err := os.Open("day_6/input_test.txt")

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	times = extract_data(scanner.Text())
	scanner.Scan()
	distances = extract_data(scanner.Text())

	for i := 0; i < len(times); i++ {
		fmt.Println(times[i], distances[i])
		mult *= get_range(times[i], distances[i])
	}

	fmt.Println(mult)
}

func get_range(time float64, distance float64) int {
	high := (- math.Sqrt(time*time - 4*distance) - time) / -2
	if math.Floor(high) != high {
		high += 1
	}
	high = math.Floor(high)
	low := (math.Sqrt(time*time - 4*distance) - time) / -2
	if math.Floor(low) == low {
		low += 1
		} else {
		low = math.Ceil(low)
	}
	fmt.Println(low, high, high-low)
	return int(high-low)
}

func extract_data(line string) []float64 {
	string_list := strings.Split(strings.Split(line, ":")[1], " ")
	var int_list []float64
	string_int := ""
	for _, s := range string_list {
		if s != "" {
			// int_list = append(int_list, string_to_int(s))
			string_int = string_int + strings.TrimSpace(s)
		}
	}
	int_list = append(int_list, string_to_int(string_int))
	return int_list
}

func string_to_int(str string) float64 {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return float64(i)
}
