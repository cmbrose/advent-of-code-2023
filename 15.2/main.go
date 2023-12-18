package main

import (
	"fmt"
	"strings"

	"main/util"
)

type lens struct {
	label       string
	focalLength int
}

func main() {
	line := util.ReadInputLines("./input.txt")[0]

	strs := strings.Split(line, ",")

	m := make([][]lens, 256)

	for _, s := range strs {
		if strings.HasSuffix(s, "-") {
			label := s[:len(s)-1]
			h := hash(label)

			lenses := m[h]
			for i := 0; i < len(lenses); i += 1 {
				l := lenses[i]
				if l.label == label {
					if i == 0 {
						m[h] = lenses[i+1:]
					} else if i == len(lenses)-1 {
						m[h] = lenses[:i]
					} else {
						m[h] = append(lenses[:i], lenses[i+1:]...)
					}
					break
				}
			}
		} else {
			pair := strings.Split(s, "=")

			l := lens{pair[0], util.AssertInt(pair[1])}
			h := hash(l.label)

			lenses := m[h]

			for i := 0; i < len(lenses); i += 1 {
				l2 := lenses[i]
				if l.label == l2.label {
					lenses[i] = l
					goto dontAppend
				}
			}

			m[h] = append(lenses, l)

		dontAppend:
		}
	}

	total := 0

	for i, lenses := range m {
		for j, l := range lenses {
			total += (1 + i) * (1 + j) * l.focalLength
		}
	}

	fmt.Printf("%d\n", total)
}

func hash(str string) int {
	var h byte

	for _, r := range str {
		h = (h + byte(r)) * 17
	}

	return int(h)
}
