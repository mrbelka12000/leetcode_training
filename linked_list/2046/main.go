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

func sortLinkedList(head *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result

	curr := head

	var check bool
	for curr != nil {
		if curr.Val < 0 {
			t := result.Next
			result.Next = &ListNode{
				Val:  curr.Val,
				Next: t,
			}
			if !check {
				tmp = result.Next
				check = true
			}

			curr = curr.Next
			continue
		}

		check = true
		tmp.Next = &ListNode{
			Val: curr.Val,
		}

		tmp = tmp.Next
		curr = curr.Next
	}

	return result.Next
}
