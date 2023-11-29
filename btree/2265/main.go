package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	var result int

	dfs(root, &result)

	return result
}

func dfs(root *TreeNode, result *int) {
	if root == nil {
		return
	}

	var sums, count int
	sum(root, &sums, &count)

	if sums/count == root.Val {
		*result++
	}

	dfs(root.Left, result)
	dfs(root.Right, result)
}

func sum(root *TreeNode, sums, count *int) {
	if root == nil {
		return
	}

	*sums += root.Val
	*count++

	sum(root.Left, sums, count)
	sum(root.Right, sums, count)
}
