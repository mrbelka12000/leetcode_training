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

func findTilt(root *TreeNode) int {
	result = 0

	runner(root)

	return result
}

var result int

func runner(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lSum, rSum := runner(root.Left), runner(root.Right)

	result += abs(lSum - rSum)

	return lSum + rSum + root.Val
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
