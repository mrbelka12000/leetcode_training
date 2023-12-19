package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		if subRoot == nil {
			return false
		}
		return true
	}
	return dfs(root, subRoot)
}

func dfs(r, sr *TreeNode) bool {
	if r == nil {
		return false
	}

	if r.Val == sr.Val {
		if runner(r, sr) {
			return true
		}
	}

	return dfs(r.Left, sr) || dfs(r.Right, sr)
}

func runner(r, sr *TreeNode) bool {
	if r == nil {
		if sr == nil {
			return true
		}
		return false
	}

	if sr == nil {
		return false
	}
	if r.Val != sr.Val {
		return false
	}
	return runner(r.Left, sr.Left) && runner(r.Right, sr.Right)
}
