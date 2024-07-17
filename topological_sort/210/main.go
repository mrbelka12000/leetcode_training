package main

import (
	"fmt"
)

func main() {
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}

func findOrder(n int, pq [][]int) []int {
	graph := make([][]int, n)

	inDegree := make([]int, n)
	for _, v := range pq {
		inDegree[v[0]]++
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	var q, result []int

	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
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

	if len(result) != n {
		return nil
	}
	return result
}
