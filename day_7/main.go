package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type hand struct {
	hand string
	amount string
	type_score int
	card_scores [5]int

}

var card_score_map = map[string]int{
	"J": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

func main() {	

	var hands []hand
	file, err := os.Open("day_7/input.txt")
	// file, err := os.Open("day_7/input_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line_split := strings.Split(line, " ")
		hands = append(hands, hand{hand: line_split[0], amount: line_split[1]})
	}
	
	for index, hand := range hands {
		hands[index] = evaluate_hand(hand)
	}
	quicksort(hands, 0, len(hands)-1)

	fmt.Println(sumHandScores(hands))
}

func sumHandScores(hands []hand) int {
	total := 0
	for i, h := range hands {
		amount, err := strconv.Atoi(h.amount)
		if err != nil {
			log.Fatal(err)
		}
		total += amount * (i + 1)
	}

	return total
}

func evaluate_hand(h hand) hand {

	var present_cards = map[string]int{
		"2": 0,
		"3": 0,
		"4": 0,
		"5": 0,
		"6": 0,
		"7": 0,
		"8": 0,
		"9": 0,
		"T": 0,
		"J": 0,
		"Q": 0,
		"K": 0,
		"A": 0,
	}

	for i, c := range h.hand {
		// fmt.Println(string(c))
		h.card_scores[i] = card_score_map[string(c)]
		present_cards[string(c)] += 1
	}

	fiveOfAKind := 0
	fourOfAKind := 0
	threeOfAKind := 0
	twoOfAKind := 0
	highCard := 1

	for k, v := range present_cards { 
		if k != "J" {
			if v == 5 {
				fiveOfAKind += 1
				break
			} else if v == 4 {
				fourOfAKind += 1
				break
			} else if v == 3 {
				threeOfAKind += 1
			} else if v == 2 {
				twoOfAKind += 1
			} else if v > 1 {
				highCard = 0
			}
		}
	}



	// 5 of a kind
	if fiveOfAKind == 1 || (fourOfAKind == 1 && present_cards["J"] == 1) || (threeOfAKind==1 && present_cards["J"] == 2) || (twoOfAKind==1 && present_cards["J"] == 3) || present_cards["J"] >= 4 {
		h.type_score = 7
	// 4 of a kind
	} else if fourOfAKind == 1 || (threeOfAKind == 1 && present_cards["J"] == 1) || (twoOfAKind == 1 && present_cards["J"] == 2) || present_cards["J"] >= 3 {
		h.type_score = 6
	// Full house
	} else if (threeOfAKind == 1 && twoOfAKind == 1) || (twoOfAKind == 2 && present_cards["J"] == 1) {
		h.type_score = 5
	// 3 of a kind
	} else if threeOfAKind == 1 || (twoOfAKind == 1 && present_cards["J"] == 1) || present_cards["J"] >= 2 {
		h.type_score = 4
	// 2 pairs
	} else if twoOfAKind == 2 || (twoOfAKind == 1 && present_cards["J"] == 1) || (present_cards["J"] == 2) {
		h.type_score = 3
	// 1 pair
	} else if twoOfAKind == 1 || present_cards["J"] == 1 {
		h.type_score = 2
	// high card
	} else if highCard == 1 {
		h.type_score = 1
	} else {
		h.type_score = 0
	}

	// if h.type_score == 7 && present_cards["J"] == 0 {
	// 	fmt.Println(h.hand, present_cards["J"], h.type_score)
	// }
	
	return h
}

func quicksort(arr []hand, left int, right int) {
	if left < right {
		partition_pos := partition(arr, left, right)
		quicksort(arr, left, partition_pos -1)
		quicksort(arr, partition_pos + 1, right)
	}
}

func swapHands(hand1 hand, hand2 hand) bool {
	if hand1.type_score < hand2.type_score {
		return true
	} else if hand1.type_score > hand2.type_score {
		return false
	} else {
		for i := 0; i < len(hand1.card_scores); i++ {
			if hand1.card_scores[i] < hand2.card_scores[i] {
				return true
			} else if hand1.card_scores[i] > hand2.card_scores[i] {
				return false 
			}
		}
	}
	return true
}

func partition(arr []hand, left int, right int) int {
	i := left
	j := right
	pivot := arr[right]

	for i < j {
		for i < right && swapHands(arr[i], pivot)  {
			i++
		}
		for j > left && swapHands(pivot, arr[j]) {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	if swapHands(pivot, arr[i]) {
		arr[i], arr[right] = arr[right], arr[i]
	}

	return i
}