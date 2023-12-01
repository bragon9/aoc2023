package day1

import (
	"aoc2023/pkg/answer"
	"aoc2023/pkg/inputreader"
	"fmt"
	"strconv"
	"unicode"
)

func Part1() (any, error) {
	lines, err := inputreader.ReadLines("pkg/days/day1/input/p1.txt")
	if err != nil {
		return nil, err
	}

	total := 0
	for _, line := range lines {
		var (
			first  rune
			second rune
		)
		for _, r := range line {
			if unicode.IsDigit(r) {
				first = r
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			r := rune(line[i])
			if unicode.IsDigit(r) {
				second = r
				break
			}
		}

		numbers, err := strconv.Atoi(string(first) + string(second))
		if err != nil {
			return nil, err
		}

		total += numbers
	}

	return total, nil
}

func Part2() (any, error) {
	lines, err := inputreader.ReadLines("pkg/days/day1/input/p2sample.txt")
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return nil, nil
}

func Solve() (answer.Answer, error) {
	var ans answer.Answer
	p1, err := Part1()
	if err != nil {
		return ans, err
	}
	p2, err := Part2()
	if err != nil {
		return ans, err
	}
	ans.Part1 = p1
	ans.Part2 = p2

	return ans, nil
}
