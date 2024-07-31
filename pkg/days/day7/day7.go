package day7

import (
	"aoc2023/pkg/answer"
	"aoc2023/pkg/inputreader"
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var cardValueMap = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

// Hand Values
const (
	HandValueNone = iota
	HandValuePair
	HandValueTwoPair
	HandValueThreeOfAKind
	HandValueFullHouse
	HandValueFourOfAKind
	HandValueFiveOfAKind
)

type Hand struct {
	Cards     [5]string
	HandScore int
	CardScore [5]int
	Bid       int
}

func checkNumPairs(cardsMap map[string]int) int {
	var numPairs int
	for _, v := range cardsMap {
		if v == 2 {
			numPairs++
		}
	}
	return numPairs
}

func scoreHand(h Hand) int {
	var max int
	var cardsMap = make(map[string]int)
	for _, card := range h.Cards {
		cardsMap[card]++
		if cardsMap[card] > max {
			max = cardsMap[card]
		}
	}

	var score int
	// Score the type of hand
	switch max {
	case 1:
		score = HandValueNone
	case 2:
		score = HandValuePair
		if numPairs := checkNumPairs(cardsMap); numPairs == 2 {
			score = HandValueTwoPair
		}
	case 3:
		score = HandValueThreeOfAKind
		if numPairs := checkNumPairs(cardsMap); numPairs == 1 {
			score = HandValueFullHouse
		}
	case 4:
		score = HandValueFourOfAKind
	case 5:
		score = HandValueFiveOfAKind
	}

	return score
}

func parseHands(lines []string) []Hand {
	var hands []Hand
	for _, line := range lines {
		var hand Hand
		split := strings.Split(line, " ")
		for j, card := range split[0] {
			hand.Cards[j] = string(card)
		}
		hand.Bid, _ = strconv.Atoi(split[1])
		hands = append(hands, hand)
	}

	return hands
}

func part1() (any, error) {
	lines, err := inputreader.ReadLines("pkg/days/day7/input/p1.txt")
	if err != nil {
		return nil, err
	}

	hands := parseHands(lines)

	for i, h := range hands {
		hands[i].HandScore = scoreHand(h)

		for j, card := range h.Cards {
			hands[i].CardScore[j] = cardValueMap[card]
		}
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		if a.HandScore == b.HandScore {
			for i := range 5 {
				if a.CardScore[i] == b.CardScore[i] {
					continue
				}
				return cmp.Compare(a.CardScore[i], b.CardScore[i])
			}
			return 0
		}
		return cmp.Compare(a.HandScore, b.HandScore)
	})

	for _, h := range hands {
		fmt.Println(h)
	}

	var score int
	for i, h := range hands {
		score += h.Bid * (i + 1)
	}

	return score, nil
}

func part2() (any, error) {
	// lines, err := inputreader.ReadLines("pkg/days/day6/input/p1.txt")
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func Solve() (answer.Answer, error) {
	return answer.Solve(part1, part2)
}
