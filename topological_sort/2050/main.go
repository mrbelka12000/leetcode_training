package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumTime(3, [][]int{{1, 3}, {1, 2}}, []int{3, 2, 5}))
}

func minimumTime(n int, relations [][]int, time []int) int {
	var (
		result   int
		inDegree = make([]int, n)
		graph    = make([][]int, n)
		calc     = make([]int, n)
		q        []int
	)

	for _, relation := range relations {
		graph[relation[0]-1] = append(graph[relation[0]-1], relation[1]-1)
		inDegree[relation[1]-1]++
	}

	for i := 1; i <= n; i++ {
		if inDegree[i-1] == 0 {
			q = append(q, i-1)
			calc[i-1] = time[i-1]
		}
	}

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			course := q[i]
			for _, edge := range graph[course] {
				inDegree[edge]--
				if inDegree[edge] == 0 {
					q = append(q, edge)
				}
				calc[edge] = max(calc[edge], calc[course]+time[edge])
			}

			result = max(result, calc[course])
		}

		q = q[size:]
	}

	return result
}
