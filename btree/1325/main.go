package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	deleteLeafs(root, target)

	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}

	return root
}

func deleteLeafs(root *TreeNode, target int) bool {

	if root == nil {
		return false
	}

	if deleteLeafs(root.Left, target) {
		root.Left = nil
	}

	if deleteLeafs(root.Right, target) {
		root.Right = nil
	}

	if root.Left == nil && root.Right == nil && root.Val == target {
		return true
	}

	return false
}
