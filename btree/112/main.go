package main

/**
 * Definition for a binary tree node.

 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return find(root, targetSum, root.Val)
}

func find(root *TreeNode, targetSum, sum int) bool {
	if root == nil {
		return false
	}

	if targetSum == sum {
		if root.Left == nil && root.Right == nil {
			return true
		}
	}

	if root.Right != nil {
		if find(root.Right, targetSum, sum+root.Right.Val) {
			return true
		}
	}

	if root.Left != nil {
		if find(root.Left, targetSum, sum+root.Left.Val) {
			return true
		}
	}

	return false
}
