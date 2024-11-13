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

func findLeaves(root *TreeNode) [][]int {
	dep := getDepth(root)

	result := make([][]int, dep)

	for i := 0; i < len(result); i++ {
		result[i] = getChilds(root, nil, false)
	}

	return result
}

func getChilds(root, parent *TreeNode, dir bool) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		if parent != nil {
			if !dir {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		}
		return []int{root.Val}
	}

	left := getChilds(root.Left, root, false)
	right := getChilds(root.Right, root, true)

	return append(left, right...)
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l, r := getDepth(root.Left), getDepth(root.Right)

	return max(l, r) + 1
}
