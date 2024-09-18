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

func numComponents(head *ListNode, nums []int) int {
	mp := make(map[int]bool)
	for _, v := range nums {
		mp[v] = true
	}

	var (
		result int
		curr   *ListNode
	)

	for head != nil {
		if mp[head.Val] {
			curr = head
		} else {
			if curr != nil {
				result++
			}
			curr = nil
		}
		head = head.Next
	}

	if curr != nil {
		result++
	}
	return result
}
