package main

import (
	"aoc2023/pkg/days/day3"
	"log"
)

func main() {
	ans, err := day3.Solve()
	if err != nil {
		log.Fatal(err)
	}

	ans.Print()
}
