package main

import (
	"fmt"
)

func main() {
	fmt.Println(gardenNoAdj(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))
}

func gardenNoAdj(n int, paths [][]int) []int {
	graph := make([][]int, n+1)
	for _, v := range paths {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	result := make([]int, n)

	var q []int
	for i := 1; i <= n; i++ {
		q = append(q, i)
	}

	flowers := make([]int, n+1)
	for len(q) != 0 {
		curr := q[0]
		q = q[1:]

		for _, v := range graph[curr] {
			if flowers[v] == 0 {
				q = append(q, v)
			}
		}

		if flowers[curr] == 0 {
			mp := map[int]bool{
				1: true, 2: true, 3: true, 4: true,
			}

			for _, v := range graph[curr] {
				delete(mp, flowers[v])
			}

			for k := range mp {
				flowers[curr] = k
				break
			}
		}
	}

	for i := 1; i < len(flowers); i++ {
		result[i-1] = flowers[i]
	}

	return result
}
