package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	graph := make(map[*TreeNode][]*TreeNode)
	leafs := make(map[*TreeNode]bool)

	var runner func(node *TreeNode)
	runner = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			leafs[node] = true
		}

		if node.Left != nil {
			graph[node] = append(graph[node], node.Left)
			graph[node.Left] = append(graph[node.Left], node)
		}

		if node.Right != nil {
			graph[node] = append(graph[node], node.Right)
			graph[node.Right] = append(graph[node.Right], node)
		}

		runner(node.Left)
		runner(node.Right)
	}
	runner(root)

	result = 0
	for k := range leafs {
		getDist(graph, k, distance, leafs, make(map[*TreeNode]bool))
	}

	return result / 2
}

var result int

func getDist(graph map[*TreeNode][]*TreeNode, curr *TreeNode, k int, leafs, visited map[*TreeNode]bool) {
	if k == -1 {
		return
	}

	if leafs[curr] && len(visited) > 0 {
		result++
	}

	visited[curr] = true

	for _, v := range graph[curr] {
		if !visited[v] {
			getDist(graph, v, k-1, leafs, visited)
		}
	}
}
