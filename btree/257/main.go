package main

import "fmt"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var result []string

	dfs(root, fmt.Sprint(root.Val), &result)

	return result
}

func dfs(root *TreeNode, tmp string, result *[]string) {
	if root.Left == nil && root.Right == nil {
		*result = append(*result, tmp)
		return
	}

	if root.Left != nil {
		dfs(root.Left, fmt.Sprintf("%v->%v", tmp, root.Left.Val), result)
	}
	if root.Right != nil {
		dfs(root.Right, fmt.Sprintf("%v->%v", tmp, root.Right.Val), result)
	}
}
