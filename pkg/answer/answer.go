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

func Solve(f1 func() (any, error), f2 func() (any, error)) (Answer, error) {
	var ans Answer
	p1, err := f1()
	if err != nil {
		return ans, err
	}
	p2, err := f2()
	if err != nil {
		return ans, err
	}
	ans.Part1 = p1
	ans.Part2 = p2

	return ans, nil
}
