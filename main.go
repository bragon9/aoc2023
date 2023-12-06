package main

import (
	"aoc2023/pkg/days/day1"
	"log"
)

func main() {
	ans, err := day1.Solve()
	if err != nil {
		log.Fatal(err)
	}

	ans.Print()
}
