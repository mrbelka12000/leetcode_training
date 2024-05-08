package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	return dfs(root, make(map[int]bool))
}

func dfs(root *TreeNode, mp map[int]bool) bool {
	if root == nil {
		return true
	}

	mp[root.Val] = true

	if len(mp) > 1 {
		return false
	}

	if !dfs(root.Left, mp) {
		return false
	}

	if !dfs(root.Right, mp) {
		return false
	}

	return true
}
