package main

import (
	"aoc2023/pkg/days/day8"
	"log"
)

func main() {
	ans, err := day8.Solve()
	if err != nil {
		log.Fatal(err)
	}

	ans.Print()
}
