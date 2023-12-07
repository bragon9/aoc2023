package day3

import (
	"aoc2023/pkg/answer"
	"aoc2023/pkg/inputreader"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var filterMap map[rune]any

func init() {
	filterMap = map[rune]any{
		'1': struct{}{},
		'2': struct{}{},
		'3': struct{}{},
		'4': struct{}{},
		'5': struct{}{},
		'6': struct{}{},
		'7': struct{}{},
		'8': struct{}{},
		'9': struct{}{},
		'0': struct{}{},
		'.': struct{}{},
	}
}

func getNumberStop(line string, start int) int {
	stop := start
	for stop < len(line) {
		r := rune(line[stop])
		if !unicode.IsDigit(r) {
			break
		}
		stop++
	}
	return stop - 1
}

func checkNumber(coords [2]string, symbolLocationMap map[string]any) bool {
	start := strings.Split(coords[0], ",")
	x1, _ := strconv.Atoi(start[0])
	y, _ := strconv.Atoi(start[1])
	x2, _ := strconv.Atoi(strings.Split(coords[1], ",")[0])

	// above
	for i := x1 - 1; i <= x2+1; i++ {
		if _, ok := symbolLocationMap[fmt.Sprintf("%v,%v", i, y-1)]; ok {
			return true
		}
	}
	// left
	if _, ok := symbolLocationMap[fmt.Sprintf("%v,%v", x1-1, y)]; ok {
		return true
	}
	// right
	if _, ok := symbolLocationMap[fmt.Sprintf("%v,%v", x2+1, y)]; ok {
		return true
	}
	// below
	for i := x1 - 1; i <= x2+1; i++ {
		if _, ok := symbolLocationMap[fmt.Sprintf("%v,%v", i, y+1)]; ok {
			return true
		}
	}

	return false
}

func part1() (any, error) {
	lines, err := inputreader.ReadLines("pkg/days/day3/input/p1.txt")
	if err != nil {
		return nil, err
	}

	numberMap := make(map[int][][2]string)
	symbolLocationMap := make(map[string]any)

	for y, line := range lines {
		x := 0
		var (
			start int
			stop  int
		)
		for x < len(line) {
			r := rune(line[x])

			// If digit, keep track
			if unicode.IsDigit(r) {
				start = x
				stop = getNumberStop(line, start)
				number, _ := strconv.Atoi(line[start : stop+1])
				numberMap[number] = append(numberMap[number], [2]string{fmt.Sprintf("%v,%v", start, y), fmt.Sprintf("%v,%v", stop, y)})
				x = stop + 1
				continue
			}
			// Store symbol locations
			if _, ok := filterMap[r]; !ok {
				symbolLocationMap[fmt.Sprintf("%v,%v", x, y)] = string(r)
			}

			x++
		}
	}

	total := 0
	for number, coords := range numberMap {
		for _, coord := range coords {
			if checkNumber(coord, symbolLocationMap) {
				total += number
			}
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
