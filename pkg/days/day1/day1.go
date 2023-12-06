package day1

import (
	"aoc2023/pkg/answer"
	"aoc2023/pkg/inputreader"
	"strconv"
	"unicode"
)

var trie map[rune]any
var reverseTrie map[rune]any

func init() {
	trie = makeTrie(false)
	reverseTrie = makeTrie(true)
}

func makeTrie(reverse bool) map[rune]any {
	words := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	tr := make(map[rune]any)
	var t map[rune]any
	for word, val := range words {
		t = tr
		if reverse {
			for i := len(word) - 1; i >= 0; i-- {
				r := rune(word[i])
				if _, ok := t[r]; !ok {
					t[r] = map[rune]any{}
				}
				t = t[r].(map[rune]any)

				// tombstone with value
				if i == 0 {
					t['.'] = val
				}
			}
		} else {
			for i, r := range word {
				if _, ok := t[r]; !ok {
					t[r] = map[rune]any{}
				}
				t = t[r].(map[rune]any)

				// tombstone with value
				if i == len(word)-1 {
					t['.'] = val
				}
			}
		}

	}

	return tr
}

func searchWord(s string, i int, reverse bool) (bool, rune) {
	stop := len(s)
	step := 1
	t := trie
	if reverse {
		stop = -1
		step = -1
		t = reverseTrie
	}

	for i != stop {
		r := rune(s[i])
		if _, ok := t[r]; !ok {
			return false, '0'
		}
		t = t[r].(map[rune]any)
		if val, found := t['.']; found {
			return true, val.(rune)
		}
		i += step
	}

	return false, '0'
}

func searchNumbers(s string, searchWords bool) (int, error) {
	var (
		first  rune
		second rune
	)

	for i, r := range s {
		if unicode.IsDigit(r) {
			first = r
			break
		}
		if searchWords {
			found, val := searchWord(s, i, false)
			if found {
				first = val
				break
			}
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])
		if unicode.IsDigit(r) {
			second = r
			break
		}
		if searchWords {
			found, val := searchWord(s, i, true)
			if found {
				second = val
				break
			}
		}
	}

	numbers, err := strconv.Atoi(string(first) + string(second))
	if err != nil {
		return 0, err
	}

	return numbers, nil
}

func Part1() (any, error) {
	lines, err := inputreader.ReadLines("pkg/days/day1/input/p1.txt")
	if err != nil {
		return nil, err
	}

	var total int
	for _, line := range lines {
		numbers, err := searchNumbers(line, false)
		if err != nil {
			return nil, err
		}

		total += numbers
	}

	return total, nil
}

func Part2() (any, error) {
	lines, err := inputreader.ReadLines("pkg/days/day1/input/p1.txt")
	if err != nil {
		return nil, err
	}

	var total int
	for _, line := range lines {
		numbers, err := searchNumbers(line, true)
		if err != nil {
			return nil, err
		}

		total += numbers
	}

	return total, nil
}

func Solve() (answer.Answer, error) {
	return answer.Solve(Part1, Part2)
}
