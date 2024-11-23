package main

import (
	"fmt"
)

func main() {
	fmt.Println(distanceToCycle(7, [][]int{{1, 2}, {2, 4}, {4, 3}, {3, 1}, {0, 1}, {5, 2}, {6, 5}}))
}

func distanceToCycle(n int, edges [][]int) []int {
	graph := make([][]int, n)

	for _, v := range edges {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	visited := make([]int, n)
	for i := 0; i < len(visited); i++ {
		visited[i] = -1
	}

	cycle := findCycle(graph, 0, -1, nil, visited)

	// fmt.Println(graph, "graph")

	return bfs(graph, cycle, n)
}

func bfs(graph [][]int, q []int, n int) []int {
	result := make([]int, n)
	visited := make([]bool, n)
	for _, v := range q {
		visited[v] = true
	}
	// fmt.Println(q)
	var dist int

	for len(q) != 0 {

		size := len(q)

		for i := 0; i < size; i++ {
			curr := q[i]
			result[curr] = dist

			for _, v := range graph[curr] {
				if !visited[v] {
					q = append(q, v)
					visited[v] = true
				}
			}
		}

		dist++
		q = q[size:]

		// fmt.Println(q)
	}

	return result
}

func findCycle(graph [][]int, curr, prev int, queue, visited []int) []int {
	if visited[curr] != -1 {
		return queue[visited[curr]:]
	}

	visited[curr] = len(queue)
	for _, v := range graph[curr] {
		if v != prev {
			if result := findCycle(graph, v, curr, append(queue, v), visited); result != nil {
				return result
			}
		}
	}

	return nil
}
