package main

import (
	"fmt"
)

func main() {
	fmt.Println(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	val := insert(&TreeNode{}, inorder, &postorder)
	return val
}

func insert(root *TreeNode, in []int, post *[]int) *TreeNode {
	if len(in) == 0 || len(*post) == 0 {
		return nil
	}

	currRoot := (*post)[len(*post)-1]
	root = &TreeNode{
		Val: currRoot,
	}
	*post = (*post)[:len(*post)-1]

	for i, v := range in {
		if v == currRoot {
			root.Right = insert(root.Left, in[i+1:], post)
			root.Left = insert(root.Left, in[:i], post)
		}
	}

	return root
}
