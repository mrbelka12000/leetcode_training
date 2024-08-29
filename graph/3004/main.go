package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumSubtreeSize([][]int{{0, 1}, {0, 2}, {0, 3}}, []int{1, 1, 2, 3}))
}

func maximumSubtreeSize(edges [][]int, colors []int) int {
	var result int
	graph := make([][]int, len(edges)+1)

	for _, v := range edges {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	visited := make([]bool, len(colors))
	var runner func(graph [][]int, curr int) (int, bool)

	runner = func(graph [][]int, curr int) (tmp int, ok bool) {
		visited[curr] = true
		currColor := colors[curr]
		var check bool
		for _, v := range graph[curr] {
			if !visited[v] {
				newSize, ok := runner(graph, v)
				if colors[v] == currColor {
					tmp += newSize
				} else {
					check = true
				}
				if !ok {
					check = true
				}
			}
		}

		if !check {
			result = max(result, tmp+1)
			return tmp + 1, true
		}

		return 1, false
	}
	val, _ := runner(graph, 0)
	return max(result, val)
}
