package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCycle([]int{3, 3, 4, 2, 3}))
}

func longestCycle(edges []int) int {
	result = 0
	visited = make([]bool, len(edges))
	for i := range edges {
		if edges[i] != -1 {
			dfs(edges, i, i, 0)
		}
	}

	if result == 0 {
		return -1
	}

	return result
}

var (
	result  int
	visited []bool
)

func dfs(edges []int, curr, start, dep int) bool {
	if curr == start && dep != 0 {
		result = max(result, dep)
		return true
	}

	if edges[curr] == -1 || visited[curr] {
		return false
	}

	visited[curr] = true
	if dfs(edges, edges[curr], start, dep+1) {
		return true
	}

	visited[curr] = false
	return false
}
