package main

import "sort"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr := &[]int{}
	dfs(root1, arr)

	dfs(root2, arr)

	sort.Ints(*arr)

	return *arr
}

func dfs(root *TreeNode, arr *[]int) []int {
	if root == nil {
		return nil
	}

	*arr = append(*arr, root.Val)
	return append(dfs(root.Left, arr), dfs(root.Right, arr)...)
}
