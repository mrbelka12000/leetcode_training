package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var result int

	dfs(root, targetSum, &result)

	return result
}

func dfs(root *TreeNode, target int, result *int) {
	if root == nil {
		return
	}
	sums := root.Val

	if sums == target {
		*result++
	}

	sum(root.Left, sums, target, result)
	sum(root.Right, sums, target, result)

	dfs(root.Left, target, result)
	dfs(root.Right, target, result)
}

func sum(root *TreeNode, sums, target int, result *int) {
	if root == nil {
		return
	}

	sums += root.Val

	if sums == target {
		*result++
	}

	sum(root.Left, sums, target, result)
	sum(root.Right, sums, target, result)
}
