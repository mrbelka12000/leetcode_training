package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(eventualSafeNodes([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}))
}

func eventualSafeNodes(edges [][]int) []int {
	inDegree := make([]int, len(edges))
	graph := make([][]int, len(edges))
	for i, v := range edges {
		inDegree[i] = len(v)
		for _, con := range v {
			graph[con] = append(graph[con], i)
		}
	}

	var result []int
	var q []int

	for i, v := range inDegree {
		if v == 0 {
			q = append(q, i)
		}
	}

	for len(q) != 0 {
		n := q[0]
		q = q[1:]
		result = append(result, n)
		for _, v := range graph[n] {
			inDegree[v]--
			if inDegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	sort.Ints(result)
	return result
}
