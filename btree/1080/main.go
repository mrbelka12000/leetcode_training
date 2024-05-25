package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	runner(root, limit, 0)

	if root.Left == nil && root.Right == nil && root.Val < limit {
		return nil
	}

	return root
}

func runner(root *TreeNode, limit, sum int) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return sum+root.Val < limit
	}

	l := runner(root.Left, limit, sum+root.Val)
	r := runner(root.Right, limit, sum+root.Val)

	if l && r {
		root.Left = nil
		root.Right = nil
		return true
	}
	if r {
		root.Right = nil
	}
	if l {
		root.Left = nil
	}

	return false
}
