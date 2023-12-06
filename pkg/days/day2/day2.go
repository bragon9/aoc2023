package day2

import (
	"aoc2023/pkg/answer"
	"aoc2023/pkg/inputreader"
	"strconv"
	"strings"
)

type Game struct {
	Number  int
	Results []map[string]int
}

func parseGames(lines []string) []Game {
	games := make([]Game, len(lines))
	for i, line := range lines {
		split := strings.Split(line, ": ")
		gameNumber, _ := strconv.Atoi(strings.Split(split[0], " ")[1])
		results := split[1]
		game := Game{
			Number: gameNumber,
		}
		for _, result := range strings.Split(results, "; ") {
			gameResults := make(map[string]int)
			for _, gameResult := range strings.Split(result, ", ") {
				draws := strings.Split(gameResult, " ")
				amount, _ := strconv.Atoi(draws[0])
				color := draws[1]
				gameResults[color] = amount
			}
			game.Results = append(game.Results, gameResults)
		}

		games[i] = game
	}

	return games
}

func checkGame(game Game, bagContents map[string]int) bool {
	for _, result := range game.Results {
		for color, resultAmount := range result {
			bagAmount, ok := bagContents[color]
			if !ok {
				return false
			}
			if bagAmount < resultAmount {
				return false
			}
		}
	}

	return true
}

func part1() (any, error) {
	lines, err := inputreader.ReadLines("pkg/days/day2/input/p1.txt")
	if err != nil {
		return 0, err
	}
	games := parseGames(lines)

	bagContents := map[string]int{
		"blue":  14,
		"red":   12,
		"green": 13,
	}

	total := 0
	for _, game := range games {
		if checkGame(game, bagContents) {
			total += game.Number
		}
	}

	return total, nil
}

func part2() (any, error) {
	return 0, nil
}

func Solve() (answer.Answer, error) {
	return answer.Solve(part1, part2)
}
