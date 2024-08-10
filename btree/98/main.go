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

func isValidBST(root *TreeNode) bool {
	return runner(root)
}

func runner(root *TreeNode) bool {
	if root == nil {
		return true
	}

	val := checker(root, root.Left, false) && checker(root, root.Right, true)
	if !val {
		return false
	}

	return runner(root.Left) && runner(root.Right)
}

func checker(root, curr *TreeNode, dir bool) bool {
	if curr == nil {
		return true
	}

	if dir && root.Val >= curr.Val {
		return false
	}

	if !dir && root.Val <= curr.Val {
		return false
	}

	return checker(root, curr.Left, dir) && checker(root, curr.Right, dir)
}
