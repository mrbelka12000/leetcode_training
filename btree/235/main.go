package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}
	var left, right *TreeNode

	if root.Val > p.Val && root.Val > q.Val {
		left = lowestCommonAncestor(root.Left, p, q)
	}

	if root.Val < p.Val && root.Val < q.Val {
		right = lowestCommonAncestor(root.Right, p, q)
	}

	if left == nil && right == nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}
