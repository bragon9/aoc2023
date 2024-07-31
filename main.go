package main

import (
	"aoc2023/pkg/days/day7"
	"log"
)

func main() {
	ans, err := day7.Solve()
	if err != nil {
		log.Fatal(err)
	}

	ans.Print()
}
