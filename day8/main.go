package day8

import (
	"github.com/nothub/AoC2023/utils"
	"log"
	"strings"
)

func init() {
	log.SetFlags(0)
}

type node struct {
	left  string
	right string
}

func level1(path string) (steps int) {
	lines := utils.ReadLines(path)

	tree := make(map[string]node)

	for _, line := range lines[1:] {
		split := strings.Split(line, " = ")
		nodes := strings.Split(strings.Trim(split[1], "()"), ", ")
		tree[split[0]] = node{left: nodes[0], right: nodes[1]}
	}

	ops := lines[0]
	cur := "AAA"

	for {
		switch ops[steps%len(ops)] {
		case 'R':
			cur = tree[cur].right
		case 'L':
			cur = tree[cur].left
		}
		steps++
		if cur == "ZZZ" {
			break
		}
	}

	return steps
}

func level2(path string) (result int) {
	_ = utils.ReadLines(path)
	return result
}
