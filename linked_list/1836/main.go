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

func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
	mp := [100001]int{}

	curr := head

	for curr != nil {
		mp[curr.Val]++
		curr = curr.Next
	}

	curr = head
	result := &ListNode{}
	tmp := result

	for curr != nil {
		if val := mp[curr.Val]; val == 1 {
			tmp.Next = &ListNode{
				Val: curr.Val,
			}
			tmp = tmp.Next
		}
		curr = curr.Next
	}

	return result.Next
}
