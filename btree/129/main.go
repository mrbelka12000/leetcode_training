package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return sumNodes(root, 0)
}

func sumNodes(node *TreeNode, num int) int {
	if node == nil {
		return 0
	}

	num = num*10 + node.Val

	if node.Left == nil && node.Right == nil {
		return num
	}

	return sumNodes(node.Left, num) + sumNodes(node.Right, num)
}
