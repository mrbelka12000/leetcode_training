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

func findDistance(root *TreeNode, p int, q int) int {
	if p == q {
		return 0
	}
	lca := LCA(root, p, q)
	l, _ := getDist(lca, q)
	r, _ := getDist(lca, p)

	return l + r
}

func getDist(root *TreeNode, p int) (int, bool) {
	if root == nil {
		return 0, false
	}
	if root.Val == p {
		return 0, true
	}

	l, okL := getDist(root.Left, p)
	if okL {
		return l + 1, true
	}
	r, okR := getDist(root.Right, p)
	if okR {
		return r + 1, true
	}

	return 0, false
}

func LCA(root *TreeNode, p, q int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p || root.Val == q {
		return root
	}

	l := LCA(root.Left, p, q)
	r := LCA(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}
	if r != nil {
		return r
	}

	return l
}
