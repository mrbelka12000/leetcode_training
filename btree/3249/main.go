package main

import (
	"fmt"
)

func main() {
	fmt.Println(countGoodNodes([][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}}))
}

func countGoodNodes(edges [][]int) int {
	graph := make([][]int, len(edges)+1)
	for _, v := range edges {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	mp = make(map[int]int)

	filler(graph, make([]bool, len(edges)+1), 0)

	result = 0

	count(graph, make([]bool, len(edges)+1), 0)
	return result
}

func count(graph [][]int, used []bool, curr int) {
	used[curr] = true
	var (
		check, canCount bool
		num             int
	)
	canCount = true
	for _, v := range graph[curr] {
		if used[v] {
			continue
		}
		check = true

		val := mp[v]

		if num == 0 {
			num = val
		}

		if num != val {
			canCount = false
		}
		count(graph, used, v)

	}
	if !check || canCount {
		result++
	}
}

var (
	mp     map[int]int
	result int
)

func filler(graph [][]int, used []bool, curr int) (result int) {
	used[curr] = true
	for _, v := range graph[curr] {
		if used[v] {
			continue
		}

		result += filler(graph, used, v)
	}

	mp[curr] = result + 1
	return result + 1
}
