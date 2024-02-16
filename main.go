package main

import (
	"aoc2023/pkg/days/day6"
	"log"
)

func main() {
	ans, err := day6.Solve()
	if err != nil {
		log.Fatal(err)
	}

	ans.Print()
}
