package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	if len(roads) == 0 {
		return 0
	}
	tree := make([][]int, len(roads)+1)
	for _, v := range roads {
		tree[v[1]] = append(tree[v[1]], v[0])
		tree[v[0]] = append(tree[v[0]], v[1])
	}
	vis := make([]bool, len(roads)+1)
	vis[0] = true
	result = 0
	dfs(tree, 0, seats, vis)
	return int64(result)
}

var result int

func dfs(tree [][]int, curr, seats int, visited []bool) int {
	tmp := 1
	for _, v := range tree[curr] {
		if !visited[v] {
			visited[v] = true
			tmp += dfs(tree, v, seats, visited)
		}
	}
	if curr != 0 {
		result += int((math.Ceil(float64(tmp) / float64(seats))))
	}
	// fmt.Printeln(tmp, seats, result)

	return tmp
}
