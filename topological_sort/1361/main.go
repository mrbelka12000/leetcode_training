package main

import (
	"fmt"
)

func main() {
	fmt.Println(validateBinaryTreeNodes(4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1}))
}

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	graph := make([][]int, n)
	inDegree := make([]int, n)
	for i, v := range leftChild {
		if v != -1 {
			graph[i] = append(graph[i], v)
			inDegree[v]++
			if inDegree[v] > 1 {
				return false
			}
		}
		r := rightChild[i]
		if r != -1 {
			graph[i] = append(graph[i], r)
			inDegree[r]++
			if inDegree[r] > 1 {
				return false
			}
		}
	}

	var q []int
	for i, v := range inDegree {
		if v == 0 {
			q = append(q, i)
			break
		}
	}

	used := make([]bool, n)
	for len(q) != 0 {
		curr := q[0]
		q = q[1:]
		used[curr] = true
		for _, v := range graph[curr] {
			inDegree[v]--
			if inDegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	for _, v := range used {
		if !v {
			return false
		}
	}
	return true
}
