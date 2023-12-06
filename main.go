package main

import (
	"aoc2023/pkg/days/day2"
	"log"
)

func main() {
	// d1, err := day1.Solve()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	d2, err := day2.Solve()
	if err != nil {
		log.Fatal(err)
	}

	d2.Print()
}
