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

func plusOne(head *ListNode) *ListNode {
	val := rec(head)
	if val == 1 {
		node := &ListNode{
			Val:  1,
			Next: head,
		}
		return node
	}

	return head
}

func rec(head *ListNode) int {
	if head == nil {
		return 0
	}
	if head.Next == nil {
		if head.Val+1 == 10 {
			head.Val = 0
			return 1
		}
		head.Val += 1
		return 0
	}

	val := rec(head.Next)

	if head.Val+val == 10 {
		head.Val = 0
		return 1
	}

	head.Val += val
	return 0
}
