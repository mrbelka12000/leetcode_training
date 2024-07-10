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

func sortedArrayToBST(nums []int) *TreeNode {
	var tree *TreeNode

	var builder func(nums []int, l, r int)
	builder = func(nums []int, l, r int) {
		if l > r {
			return
		}
		mid := (l + r) / 2
		tree = insert(tree, nums[mid])

		builder(nums, mid+1, r)
		builder(nums, l, mid-1)
	}

	builder(nums, 0, len(nums)-1)

	return tree
}

func insert(t *TreeNode, val int) *TreeNode {
	if t == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if t.Val > val {
		t.Left = insert(t.Left, val)
	} else {
		t.Right = insert(t.Right, val)
	}

	return t
}
