package day8

import (
	"aoc2023/pkg/answer"
	"aoc2023/pkg/inputreader"
	"strings"
)

type Node struct {
	Name      string
	LeftNode  *Node
	RightNode *Node
}

type Map struct {
	Directions []string
	Nodes      map[string]*Node
}

func NewNode(name string) *Node {
	return &Node{
		Name: name,
	}
}

func (m *Map) Traverse() int {
	var ptr int
	curr := m.Nodes["AAA"]
	for {
		if curr.Name == "ZZZ" {
			break
		}
		d := m.Directions[ptr%len(m.Directions)]
		if d == "L" {
			curr = curr.LeftNode
		} else {
			curr = curr.RightNode
		}

		ptr++
	}

	return ptr
}

func parseMap(lines []string) *Map {
	var m Map
	m.Nodes = make(map[string]*Node)

	directions := make([]string, 0)
	for _, r := range lines[0] {
		directions = append(directions, string(r))
	}
	m.Directions = directions

	for _, line := range lines[2:] {
		split := strings.Split(line, " = ")
		name := split[0]
		connections := strings.Split(strings.TrimFunc(split[1], func(r rune) bool {
			return r == '(' || r == ')'
		}), ", ")
		leftNode := connections[0]
		rightNode := connections[1]

		if _, ok := m.Nodes[name]; !ok {
			m.Nodes[name] = NewNode(name)
		}

		if _, ok := m.Nodes[leftNode]; !ok {
			m.Nodes[leftNode] = NewNode(leftNode)
		}

		if _, ok := m.Nodes[rightNode]; !ok {
			m.Nodes[rightNode] = NewNode(rightNode)
		}

		m.Nodes[name].LeftNode = m.Nodes[leftNode]
		m.Nodes[name].RightNode = m.Nodes[rightNode]
	}

	return &m
}

func part1() (any, error) {
	lines, err := inputreader.ReadLines("pkg/days/day8/input/input.txt")
	if err != nil {
		return nil, err
	}

	m := parseMap(lines)
	numSteps := m.Traverse()

	return numSteps, nil
}

func part2() (any, error) {
	// lines, err := inputreader.ReadLines("pkg/days/day8/input/sample.txt")
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func Solve() (answer.Answer, error) {
	return answer.Solve(part1, part2)
}
