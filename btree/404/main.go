package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	return runner(root, false)
}

func runner(root *TreeNode, dir bool) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil && dir {
		return root.Val
	}
	return runner(root.Left, true) + runner(root.Right, false)
}
