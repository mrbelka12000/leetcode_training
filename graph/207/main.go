package main

import (
	"fmt"
)

func main() {
	fmt.Println(canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	for _, pre := range prerequisites {
		x := pre[0]
		y := pre[1]
		graph[y] = append(graph[y], x)
	}

	//fmt.Println(graph)

	visited := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if hasCycle(graph, visited, i) {
			return false
		}
		//fmt.Println(visited, i)
	}

	return true
}

func hasCycle(graph [][]int, visited []int, node int) bool {
	if visited[node] == 1 {
		return true
	}
	if visited[node] == -1 {
		return false
	}

	visited[node] = 1

	for _, neighbor := range graph[node] {
		if hasCycle(graph, visited, neighbor) {
			return true
		}
	}

	visited[node] = -1
	return false
}
