package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var answer int = -123421

	solver(root, &answer)

	return answer
}

func solver(root *TreeNode, answer *int) int {
	if root == nil {
		return 0
	}

	leftMax := max(solver(root.Left, answer), 0)
	rightMax := max(solver(root.Right, answer), 0)
	*answer = max(leftMax+rightMax+root.Val, *answer)

	return max(leftMax, rightMax) + root.Val
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
