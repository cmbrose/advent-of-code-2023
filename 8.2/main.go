package main

import (
	"fmt"
	"regexp"
	"strings"

	"main/util"
)

type node struct {
	name        string
	left, right *node
}

func main() {
	lines := util.ReadInputLines("./input.txt")
	seq := lines[0]

	re := regexp.MustCompile(`([A-Z0-9]+) = \(([A-Z0-9]+), ([A-Z0-9]+)\)`)
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

	ns := util.Where(util.Values(nodes), func(n *node) bool {
		return strings.HasSuffix(n.name, "A")
	})

	loopSizes := util.Map(ns, func(n *node) int {
		for i := 0; ; i += 1 {
			if strings.HasSuffix(n.name, "Z") {
				return i
			}

			dir := seq[i%len(seq)]

			if dir == 'L' {
				n = n.left
			} else {
				n = n.right
			}
		}
	})

	fmt.Println(util.LCM(loopSizes...))
}
