package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumSemesters(3, [][]int{{1, 3}, {2, 3}}))
}

func minimumSemesters(n int, relations [][]int) int {
	graph := make([][]int, n)
	ts := make([]int, n)
	for _, v := range relations {
		graph[v[0]-1] = append(graph[v[0]-1], v[1]-1)
		ts[v[1]-1]++
	}

	var q []int

	for i, v := range ts {
		if v == 0 {
			q = append(q, i)
		}
	}
	if len(q) == 0 {
		return -1
	}

	var result int
	for len(q) != 0 {
		size := len(q)
		result++
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]
			for _, v := range graph[n] {
				ts[v]--
				if ts[v] == 0 {
					q = append(q, v)
				}
			}
		}
	}

	for _, v := range ts {
		if v > 0 {
			return -1
		}
	}

	return result
}
