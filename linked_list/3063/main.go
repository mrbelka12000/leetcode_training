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

func frequenciesOfElements(head *ListNode) *ListNode {
	arr := [100001]int{}
	for i := 0; i < len(arr); i++ {
		arr[i] = 0
	}

	curr := head

	for curr != nil {
		arr[curr.Val]++
		curr = curr.Next
	}

	result := &ListNode{}
	tmp := result
	curr = head

	for curr != nil {
		if val := arr[curr.Val]; val != 0 {
			tmp.Next = &ListNode{
				Val: val,
			}
			tmp = tmp.Next
			// delete(mp, curr.Val)
			arr[curr.Val] = 0
		}

		curr = curr.Next
	}

	return result.Next
}
