package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, result := runner(root, 0)

	return result
}

func runner(root *TreeNode, dep int) (int, bool) {
	if root == nil {
		return dep, true
	}

	left, lok := runner(root.Left, dep+1)
	right, rok := runner(root.Right, dep+1)

	if !lok || !rok {
		return 0, false
	}

	if abs(left-right) > 1 {
		return 0, false
	}

	return max(left, right), true
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
