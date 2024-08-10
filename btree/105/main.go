package main

import (
	"fmt"
)

func main() {
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	root := &TreeNode{
		Val: preorder[0],
	}
	runner(preorder, inorder, "root", root)
	return root
}

func runner(pre, in []int, dir string, root *TreeNode) *TreeNode {
	if len(pre) == 0 || len(in) == 0 {
		return nil
	}

	if len(in) == 1 {
		return &TreeNode{
			Val: in[0],
		}
	}
	val := pre[0]

	root.Val = val

	for i := 0; i < len(in); i++ {
		if in[i] == val {
			root.Left = runner(pre[1:], in[:i], "left", &TreeNode{})
			root.Right = runner(pre[i+1:], in[i+1:], "right", &TreeNode{})
		}
	}

	return root
}
