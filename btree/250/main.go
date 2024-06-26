package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countUnivalSubtrees(root *TreeNode) int {
	var (
		dfs    func(node *TreeNode) map[int]bool
		result int
	)

	dfs = func(node *TreeNode) map[int]bool {
		if node == nil {
			return nil
		}

		mpL := dfs(node.Left)
		mpR := dfs(node.Right)
		mp := map[int]bool{
			node.Val: true,
		}

		for k, v := range mpL {
			mp[k] = v
		}

		for k, v := range mpR {
			mp[k] = v
		}

		if len(mp) == 1 {
			result++
		}

		// fmt.Println(mp, node)
		return mp
	}

	dfs(root)

	return result
}
