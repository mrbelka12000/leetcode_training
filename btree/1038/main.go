package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	totalSum = 0

	converter(root)

	return root
}

var totalSum int

func converter(root *TreeNode) {
	if root == nil {
		return
	}

	converter(root.Right)
	root.Val += totalSum
	totalSum = root.Val
	converter(root.Left)
}
