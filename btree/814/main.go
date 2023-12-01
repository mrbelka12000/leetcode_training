package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {

	if !contain(root) {
		return nil
	}

	dfs(root)

	return root
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	if !contain(root.Left) {
		root.Left = nil
	}

	if !contain(root.Right) {
		root.Right = nil
	}

	dfs(root.Left)
	dfs(root.Right)
}

func contain(root *TreeNode) bool {
	if root == nil {
		return false
	}
	q := []*TreeNode{root}

	for len(q) != 0 {

		n := q[0]
		q = q[1:]

		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}

		if n.Val == 1 {
			return true
		}
	}

	return false
}
