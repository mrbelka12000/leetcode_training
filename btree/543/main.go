package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var result int
	runner(root, &result)

	return result
}

func runner(root *TreeNode, depth *int) {
	if root == nil {
		return
	}
	var left, right int
	dfs(root.Left, &left, 1)
	dfs(root.Right, &right, 1)

	if *depth < left+right {
		*depth = left + right
	}

	runner(root.Left, depth)
	runner(root.Right, depth)
}

func dfs(root *TreeNode, depth *int, curr int) {
	if root == nil {
		return
	}

	if curr > *depth {
		*depth = curr
	}

	dfs(root.Left, depth, curr+1)
	dfs(root.Right, depth, curr+1)
}
