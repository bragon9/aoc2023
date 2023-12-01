package answer

import "fmt"

type Answer struct {
	Part1 any
	Part2 any
}

func (a Answer) Print() {
	fmt.Println("Part 1:", a.Part1)
	fmt.Println("Part 2:", a.Part2)
}
