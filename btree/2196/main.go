package main

import (
	"fmt"
)

func main() {
	fmt.Println(createBinaryTree([][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(d [][]int) *TreeNode {
	tree := make(map[int]*TreeNode)

	mp := make(map[int]bool)
	for _, v := range d {
		parent, ok := tree[v[0]]
		if !ok {
			parent = &TreeNode{
				Val: v[0],
			}
			tree[v[0]] = parent
		}

		node, ok := tree[v[1]]
		if !ok {
			node = &TreeNode{
				Val: v[1],
			}
			tree[v[1]] = node
		}

		if v[2] == 1 {
			parent.Left = node
		} else {
			parent.Right = node
		}
		mp[v[1]] = true
	}

	for _, v := range d {
		if !mp[v[0]] {
			return tree[v[0]]
		}
	}

	return nil
}
