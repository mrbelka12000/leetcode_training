package main

import (
	"fmt"
)

func main() {
	fmt.Println(killProcess([]int{1, 3, 10, 5}, []int{3, 0, 5, 3}, 5))
}

func killProcess(pid []int, ppid []int, kill int) []int {
	graph := make(map[int][]int)
	n := len(pid)
	for i := 0; i < n; i++ {
		graph[ppid[i]] = append(graph[ppid[i]], pid[i])
	}

	var result []int

	q := []int{kill}

	for len(q) != 0 {
		curr := q[0]
		q = q[1:]
		result = append(result, curr)

		for _, neig := range graph[curr] {
			q = append(q, neig)
		}
	}

	return result
}
