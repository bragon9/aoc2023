package day4

import (
	"aoc2023/pkg/answer"
	"aoc2023/pkg/inputreader"
	"strconv"
	"strings"
)

type Card struct {
	CardNum int
	Winners map[int]struct{}
	Numbers []int
	Matches int
	Score   int
}

func (c *Card) CountScore() {
	score := 0
	for _, number := range c.Numbers {
		if _, ok := c.Winners[number]; ok {
			c.Matches++
			switch score {
			case 0:
				score = 1
			default:
				score *= 2
			}
		}
	}
	c.Score = score
}

func parseCards(lines []string) []Card {
	cards := make([]Card, 0)
	for _, line := range lines {
		split := strings.Split(line, ": ")
		cardNum, _ := strconv.Atoi(strings.Split(split[0], " ")[1])
		numbersStr := strings.Split(split[1], " | ")
		winnersArr := strings.Fields(numbersStr[0])
		numbersArr := strings.Fields(numbersStr[1])

		winners := make(map[int]struct{}, 0)
		for _, winner := range winnersArr {
			intWinner, _ := strconv.Atoi(winner)
			winners[intWinner] = struct{}{}
		}

		numbers := []int{}
		for _, number := range numbersArr {
			intNumber, _ := strconv.Atoi(number)
			numbers = append(numbers, intNumber)
		}

		card := Card{
			CardNum: cardNum,
			Winners: winners,
			Numbers: numbers,
		}
		card.CountScore()
		cards = append(cards, card)
	}

	return cards
}

func part1() (any, error) {
	lines, err := inputreader.ReadLines("pkg/days/day4/input/p1.txt")
	if err != nil {
		return nil, err
	}

	cards := parseCards(lines)
	total := 0
	for _, card := range cards {
		total += card.Score
	}

	return total, nil
}

func part2() (any, error) {
	lines, err := inputreader.ReadLines("pkg/days/day4/input/p1.txt")
	if err != nil {
		return nil, err
	}

	cards := parseCards(lines)

	total := 0
	extra_copies := map[int]int{}
	for i, card := range cards {
		total += extra_copies[i] + 1
		for j := 0; j < card.Matches; j++ {
			extra_copies[i+j+1] += extra_copies[i] + 1
		}
	}

	return total, nil
}

func Solve() (answer.Answer, error) {
	return answer.Solve(part1, part2)
}
