package main

import (
	"fmt"
)

func main() {
	fmt.Println(getAncestors(8, [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}))
}

func getAncestors(n int, edges [][]int) [][]int {
	result := make([][]int, n)

	inDegree := make([]int, n)
	graph := make([][]int, n)

	for _, v := range edges {
		graph[v[0]] = append(graph[v[0]], v[1])
		inDegree[v[1]]++
	}

	var q []int

	for i, v := range inDegree {
		if v == 0 {
			q = append(q, i)
		}
	}

	check := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		check[i] = make(map[int]bool)
	}

	for len(q) != 0 {
		n := q[0]
		q = q[1:]
		for _, v := range graph[n] {
			check[v][n] = true
			for k := range check[n] {
				check[v][k] = true
			}

			inDegree[v]--
			if inDegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	for i, mp := range check {
		for k := range mp {
			result[i] = append(result[i], k)
		}
		sort.Ints(result[i])
	}

	return result
}
