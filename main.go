package main

import (
	"aoc2023/pkg/days/day4"
	"log"
)

func main() {
	ans, err := day4.Solve()
	if err != nil {
		log.Fatal(err)
	}

	ans.Print()
}
