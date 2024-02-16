package day6

import (
	"aoc2023/pkg/answer"
	"aoc2023/pkg/inputreader"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func parseLinesV1(lines []string) []Race {
	races := []Race{}
	times := strings.Fields(strings.Split(lines[0], ":")[1])
	distances := strings.Fields(strings.Split(lines[1], ":")[1])

	for i := range len(times) {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		races = append(races, Race{Time: time, Distance: distance})
	}

	return races
}

func parseLinesV2(lines []string) Race {
	time := strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", "")
	intTime, _ := strconv.Atoi(time)
	distance := strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", "")
	intDistance, _ := strconv.Atoi(distance)

	return Race{Time: intTime, Distance: intDistance}
}

func GetWinningCombinations(r Race) int {
	answers := 0

	for j := 1; j < r.Time; j++ {
		if j*(r.Time-j) > r.Distance {
			answers += 1
		}
	}

	return answers
}

func GetTotalWinningCombinations(races []Race) int {
	total := 1
	for i := range len(races) {
		total *= GetWinningCombinations(races[i])
	}

	return total
}

func part1() (any, error) {
	lines, err := inputreader.ReadLines("pkg/days/day6/input/p1.txt")
	if err != nil {
		return nil, err
	}

	races := parseLinesV1(lines)

	return GetTotalWinningCombinations(races), nil
}

func part2() (any, error) {
	lines, err := inputreader.ReadLines("pkg/days/day6/input/p1.txt")
	if err != nil {
		return nil, err
	}

	race := parseLinesV2(lines)

	return GetWinningCombinations(race), nil
}

func Solve() (answer.Answer, error) {
	return answer.Solve(part1, part2)
}
