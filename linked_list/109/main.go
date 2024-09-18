package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	var runner func(head *ListNode)

	var result *TreeNode
	runner = func(node *ListNode) {
		if node == nil {
			return
		}
		if node.Next == nil {
			result = insert(result, node.Val)
			return
		}
		var prev *ListNode
		fast, slow := node, node
		for fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			prev = slow
			slow = slow.Next
		}
		if prev != nil {
			prev.Next = nil
		}
		result = insert(result, slow.Val)
		runner(node)
		runner(slow.Next)
	}

	runner(head)

	return result
}

func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if root.Val > val {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}
	return root
}
