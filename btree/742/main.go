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

func findClosestLeaf(root *TreeNode, k int) int {
	graph := [1001][]int{}
	leaf := [1001]bool{}
	q := []*TreeNode{root}

	for len(q) != 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			if n.Left == nil && n.Right == nil {
				leaf[n.Val] = true
			}
			if n.Left != nil {
				graph[n.Val] = append(graph[n.Val], n.Left.Val)
				graph[n.Left.Val] = append(graph[n.Left.Val], n.Val)
				q = append(q, n.Left)
			}

			if n.Right != nil {
				graph[n.Val] = append(graph[n.Val], n.Right.Val)
				graph[n.Right.Val] = append(graph[n.Right.Val], n.Val)
				q = append(q, n.Right)
			}
		}
	}

	iq := []int{k}
	var visited [10001]bool
	for len(iq) != 0 {
		n := iq[0]
		iq = iq[1:]
		if leaf[n] {
			return n
		}

		for _, v := range graph[n] {
			if !visited[v] {
				iq = append(iq, v)
				visited[v] = true
			}
		}
	}

	return k
}
