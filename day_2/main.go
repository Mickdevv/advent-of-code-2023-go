package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type game struct {
	red int
	green int
	blue int
}

func main() {
	// gamesMap := map[string]game{}

	idSum := 0
	total_power := 0

    file, err := os.Open("day_2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		game := game{red: 0, green: 0, blue: 0}
		line := scanner.Text()

		gameId := strings.Split(strings.Split(line, ":")[0], " ")[1]

		rounds := strings.Split(strings.Split(line, ":")[1], ";")

		for _, round := range rounds {
			values := strings.Split(strings.TrimSpace(round), ", ")
			for _, value := range values {
				v, err := strconv.Atoi(strings.Split(value, " ")[0])
				if err != nil {
					log.Fatal(err)
				}
				colour := strings.Split(value, " ")[1]
				if err != nil {
					log.Fatal(err)
				}
				if colour == "red" && v > game.red {
					game.red = v
				} else if colour == "green" && v > game.green {
					game.green = v
				} else if colour == "blue" && v > game.blue {
					game.blue = v
				}
			}
		}

		if game.red <= 12 && game.green <= 13 && game.blue <= 14 {
			gameId_int, _ := strconv.Atoi(gameId)

			idSum += gameId_int
		}
		total_power += game.red * game.green * game.blue

		fmt.Println(gameId, game, idSum, total_power)
	}
}