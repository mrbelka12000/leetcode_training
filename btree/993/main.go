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

func isCousins(root *TreeNode, x int, y int) bool {
	rx, depx := getDepth(root, x)
	ry, depy := getDepth(root, y)

	if rx == ry {
		return false
	}

	return depx == depy
}

func getDepth(root *TreeNode, x int) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}

	if root.Left != nil {
		if root.Left.Val == x {
			return root, 1
		}
	}

	if root.Right != nil {
		if root.Right.Val == x {
			return root, 1
		}
	}

	n, dep := getDepth(root.Left, x)
	if n != nil {
		return n, dep + 1
	}

	n1, dep1 := getDepth(root.Right, x)
	if n1 != nil {
		return n1, dep1 + 1
	}
	return nil, 0
}
