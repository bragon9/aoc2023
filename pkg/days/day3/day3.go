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

func checkNumber(number int, coords [2]string, symbolLocationMap map[string][]int) bool {
	start := strings.Split(coords[0], ",")
	x1, _ := strconv.Atoi(start[0])
	y, _ := strconv.Atoi(start[1])
	x2, _ := strconv.Atoi(strings.Split(coords[1], ",")[0])

	// above
	for i := x1 - 1; i <= x2+1; i++ {
		if _, ok := symbolLocationMap[fmt.Sprintf("%v,%v", i, y-1)]; ok {
			symbolLocationMap[fmt.Sprintf("%v,%v", i, y-1)] = append(symbolLocationMap[fmt.Sprintf("%v,%v", i, y-1)], number)
			return true
		}
	}
	// left
	if _, ok := symbolLocationMap[fmt.Sprintf("%v,%v", x1-1, y)]; ok {
		symbolLocationMap[fmt.Sprintf("%v,%v", x1-1, y)] = append(symbolLocationMap[fmt.Sprintf("%v,%v", x1-1, y)], number)
		return true
	}
	// right
	if _, ok := symbolLocationMap[fmt.Sprintf("%v,%v", x2+1, y)]; ok {
		symbolLocationMap[fmt.Sprintf("%v,%v", x2+1, y)] = append(symbolLocationMap[fmt.Sprintf("%v,%v", x2+1, y)], number)
		return true
	}
	// below
	for i := x1 - 1; i <= x2+1; i++ {
		if _, ok := symbolLocationMap[fmt.Sprintf("%v,%v", i, y+1)]; ok {
			symbolLocationMap[fmt.Sprintf("%v,%v", i, y+1)] = append(symbolLocationMap[fmt.Sprintf("%v,%v", i, y+1)], number)
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
	symbolLocationMap := make(map[string][]int)

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
				symbolLocationMap[fmt.Sprintf("%v,%v", x, y)] = []int{}
			}

			x++
		}
	}

	total := 0
	for number, coords := range numberMap {
		for _, coord := range coords {
			if checkNumber(number, coord, symbolLocationMap) {
				total += number
			}
		}
	}

	return total, nil
}

func part2() (any, error) {
	lines, err := inputreader.ReadLines("pkg/days/day3/input/p1.txt")
	if err != nil {
		return nil, err
	}

	numberMap := make(map[int][][2]string)
	gearLocationMap := make(map[string][]int)

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
			if r == '*' {
				gearLocationMap[fmt.Sprintf("%v,%v", x, y)] = []int{}
			}

			x++
		}
	}

	for number, coords := range numberMap {
		for _, coord := range coords {
			// Loop through and update gearLocationMap values
			checkNumber(number, coord, gearLocationMap)
		}
	}

	total := 0
	for _, numbers := range gearLocationMap {
		if len(numbers) == 2 {
			total += numbers[0] * numbers[1]
		}
	}

	return total, nil
}

func Solve() (answer.Answer, error) {
	return answer.Solve(part1, part2)
}
