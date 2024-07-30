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

func deleteNodes(head *ListNode, m int, n int) *ListNode {
	mm, nn := m, n
	res := head

	for head != nil {
		if mm == 1 {
			curr := head
			for curr != nil && nn != -1 {
				curr = curr.Next
				nn--
			}
			head.Next = curr
			head = head.Next
			nn = n
			mm = m
			continue
		}
		mm--
		head = head.Next
	}

	return res
}
