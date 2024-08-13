package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
}

func minMutation(startGene string, endGene string, bank []string) int {
	q := []info{{val: startGene, count: 0}}
	calc := make(map[string]int)

	for _, v := range bank {
		calc[v] = math.MaxInt32
	}
	calc[endGene] = math.MaxInt32

	for len(q) != 0 {

		size := len(q)

		for i := 0; i < size; i++ {
			n := q[i]

			g := n.val
			count := n.count

			for _, v := range bank {

				newC := getDist(g, v)
				if newC != 1 {
					continue
				}
				if calc[v] > newC+count {
					calc[v] = newC + count
					q = append(q, info{val: v, count: calc[v]})
				}
			}
		}

		q = q[size:]
	}

	if calc[endGene] == math.MaxInt32 {
		return -1
	}

	return calc[endGene]
}

type info struct {
	val   string
	count int
}

func getDist(n1, n2 string) (result int) {

	for i := 0; i < len(n1); i++ {
		if n1[i] != n2[i] {
			result++
		}
	}

	return result
}
