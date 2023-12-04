package main

import "fmt"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	mp := make(map[string]*ListNode)

	curr := head

	for curr != nil {

		if val, ok := mp[fmt.Sprint(curr)]; ok {
			// return val
			fmt.Println("map", val)
			return val
			_ = val
		}

		// fmt.Println(curr.Val, ind)

		mp[fmt.Sprint(curr)] = curr

		curr = curr.Next
	}

	return nil
}
