package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type Node struct {
	Val      int
	Children []*Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	root *Node
}

func Constructor() *Codec {
	return &Codec{}
}

func (this *Codec) encode(root *Node) *TreeNode {
	if root == nil {
		return nil
	}
	this.root = root

	q := []*Node{root}
	var result *TreeNode
	for len(q) != 0 {
		n := q[0]
		q = q[1:]
		for _, v := range n.Children {
			result = insert(result, v.Val)
			q = append(q, v)
		}
	}

	return result
}

func (this *Codec) decode(root *TreeNode) *Node {
	return this.root
}

func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if root.Left == nil {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * bst := obj.encode(root);
 * ans := obj.decode(bst);
 */
