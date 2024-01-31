package main

import (
	"aoc2023/pkg/days/day5"
	"log"
)

func main() {
	ans, err := day5.Solve()
	if err != nil {
		log.Fatal(err)
	}

	ans.Print()
}
