package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	result = 0

	runner(root)

	return result
}

var result int

func runner(root *TreeNode) (int, int) {
	if root == nil {
		return 1001, 0
	}

	lval, lcount := runner(root.Left)
	rval, rcount := runner(root.Right)

	switch {
	case lval == root.Val && rval == root.Val:
		result = max(result, lcount+rcount+2)
		return root.Val, max(lcount, rcount) + 1
	case lval == root.Val:
		result = max(result, lcount+1)
		return root.Val, lcount + 1
	case rval == root.Val:
		result = max(result, rcount+1)
		return root.Val, rcount + 1
	default:
		return root.Val, 0
	}
}
