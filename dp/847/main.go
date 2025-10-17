package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
}

func shortestPathLength(graph [][]int) int {
	var (
		result int = math.MaxInt32
		runner func(curr, cost int) int
	)

	for i := 0; i < len(graph); i++ {
		visited := make([]bool, len(graph))
		runner = func(curr, cost int) int {
			if visited[curr] {
				return cost
			}

			visited[curr] = true

			var totalCost int
			for _, v := range graph[curr] {
				runner(v, cost+1)
			}

			return totalCost + 1
		}

		result = min(result, runner(i, 0))
	}

	return result
}
