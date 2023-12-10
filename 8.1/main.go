package main

import (
	"fmt"
	"regexp"

	"main/util"
)

type node struct {
	name        string
	left, right *node
}

func main() {
	lines := util.ReadInputLines("./input.txt")
	seq := lines[0]

	re := regexp.MustCompile(`([A-Z]+) = \(([A-Z]+), ([A-Z]+)\)`)
	nodes := make(map[string]*node)

	getOrCreateNode := func(name string) *node {
		if n, ok := nodes[name]; ok {
			return n
		}

		n := &node{name: name}
		nodes[name] = n
		return n
	}

	for _, line := range lines[2:] {
		m := re.FindStringSubmatch(line)
		n, l, r := getOrCreateNode(m[1]), getOrCreateNode(m[2]), getOrCreateNode(m[3])

		n.left = l
		n.right = r
	}

	n := nodes["AAA"]
	i := 0
	for ; n.name != "ZZZ"; i += 1 {
		dir := seq[i%len(seq)]
		if dir == 'L' {
			n = n.left
		} else {
			n = n.right
		}
	}

	fmt.Printf("%d\n", i)
}
