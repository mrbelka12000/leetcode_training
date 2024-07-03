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

func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	if root1 == nil {
		return false
	}

	if find(root2, target-root1.Val) {
		return true
	}

	return twoSumBSTs(root1.Left, root2, target) || twoSumBSTs(root1.Right, root2, target)
}

func find(root2 *TreeNode, target int) bool {
	if root2 == nil {
		return false
	}
	if root2.Val == target {
		return true
	}

	if root2.Val > target {
		return find(root2.Left, target)
	}

	return find(root2.Right, target)
}
