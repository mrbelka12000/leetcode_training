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

func equalToDescendants(root *TreeNode) int {
	result = 0

	runner(root)

	return result
}

var result int

func runner(root *TreeNode) int {
	if root == nil {
		return 0
	}

	val := runner(root.Left) + runner(root.Right)
	if val == root.Val {
		result++
	}

	return root.Val + val
}
