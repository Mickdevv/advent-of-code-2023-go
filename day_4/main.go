package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type card struct {
	id int
	winning_numbers []int
	numbers []int
	matches int
	score int
}

func main() {
	file, err := os.Open("day_4/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	totalScore := 0
	totalCardCount := 0
	var original_cards []card

	cardIndex := 1

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		cardId, err := strconv.Atoi(strings.Split(strings.Split(line, ": ")[0], " ")[len(strings.Split(strings.Split(line, ": ")[0], " ")) -1])
		if err != nil {
			fmt.Println(err)
		}
		cardIndex += 1
		winning_numbers := stringArrayToIntArray(strings.Split(strings.Split(strings.Split(line, ": ")[1], " | ")[0], " "))
		card_numbers := stringArrayToIntArray(strings.Split(strings.Split(strings.Split(line, ": ")[1], " | ")[1], " "))
		
		card := card{ id: cardId, winning_numbers: winning_numbers, numbers: card_numbers }
		card_score, card_matches := calculateScore(card)
		card.score = card_score
		card.matches = card_matches
		
		original_cards = append(original_cards, card)
		totalScore += card.score
		
	}
	totalCardCount += len(original_cards)

	processing_queue := original_cards

	for len(processing_queue) > 0 {
		current_card := processing_queue[0]
		if len(processing_queue) > 1 {
			processing_queue = processing_queue[1:]
		} else {
			processing_queue = processing_queue[:0]
		}
		for i := range current_card.matches {
			processing_queue = append(processing_queue, original_cards[current_card.id + i])
			totalCardCount += 1
		}
	}

	fmt.Println(totalScore)
	fmt.Println(totalCardCount)
}

func stringArrayToIntArray(arr []string) (intArr []int) {
	for _, item := range arr {
		// fmt.Println(string(item))
		if item != "" {
			intItem, err := strconv.Atoi(strings.TrimSpace(item))
			if err != nil {
				fmt.Println(err)
			}
			intArr = append(intArr, intItem)
		}
	}
	return 
}

func calculateScore(card card) (score int, matches int) {
	score = 0
	matches = 0
	for _, i := range card.numbers {
		for _, j := range card.winning_numbers {
			if i == j {
				// fmt.Println(i, j)
				score = score * 2
				matches += 1
				if score == 0 {
					score = 1
				}
			}
		}
	}
	// score = score / 2
	return 
}