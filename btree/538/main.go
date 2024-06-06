package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	totalSum = getSum(root)
	converter(root)
	return root
}

var totalSum int

func converter(root *TreeNode) {
	if root == nil {
		return
	}

	converter(root.Left)
	root.Val, totalSum = totalSum, totalSum-root.Val
	converter(root.Right)
}

func getSum(root *TreeNode) (sum int) {
	if root == nil {
		return 0
	}

	sum += getSum(root.Left)
	sum += getSum(root.Right)

	return sum + root.Val
}
