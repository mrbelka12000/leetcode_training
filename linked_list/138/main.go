package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	mp := make(map[*Node]int)
	var arr []*Node
	curr := head
	result := &Node{}
	tmp := result

	var i int
	for curr != nil {
		mp[curr] = i
		tmp.Next = &Node{
			Val: curr.Val,
		}
		tmp = tmp.Next
		arr = append(arr, tmp)
		curr = curr.Next
		i++
	}

	curr = head
	tmp = result.Next
	for curr != nil {
		check := curr
		curr = curr.Next
		random := check.Random
		if random != nil {
			tmp.Random = arr[mp[random]]
		}
		tmp = tmp.Next
	}

	return result.Next
}
