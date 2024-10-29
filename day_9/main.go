package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("day_9/input_test.txt")

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	fmt.Println(scanner.Text())
}

func P1() {

}

func P2() {
	
}