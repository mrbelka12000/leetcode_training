package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
}

type Node struct {
	Val  int
	Next *Node
}

func insert(head *Node, x int) *Node {
	if head == nil {
		n := &Node{
			Val: x,
		}
		n.Next = n
		return n
	}

	if head.Next == head {
		n := &Node{
			Val:  x,
			Next: head,
		}
		head.Next = n
		return head
	}

	maxVal, minVal := getValues(head)

	if x <= minVal {
		return insertBefore(head, x, minVal)
	}

	if x >= maxVal {
		return insertAfter(head, x, maxVal)
	}

	mp := make(map[*Node]bool)
	curr := head
	for !mp[curr] {
		if curr.Val <= x && curr.Next.Val >= x {
			n := curr.Next
			curr.Next = &Node{
				Val:  x,
				Next: n,
			}
			return head
		}

		mp[curr] = true
		curr = curr.Next
	}

	return head
}

func insertAfter(head *Node, x, val int) *Node {
	mp := make(map[*Node]bool)

	curr := head

	for !mp[curr] {
		if curr.Val == val && curr.Next.Val != val {
			n := curr.Next

			curr.Next = &Node{
				Val:  x,
				Next: n,
			}
			return head
		}
		if mp[curr.Next] {
			n := curr.Next
			curr.Next = &Node{
				Val:  x,
				Next: n,
			}
			return head
		}

		mp[curr] = true
		curr = curr.Next
	}
	return head
}

func insertBefore(head *Node, x, val int) *Node {
	mp := make(map[*Node]bool)

	curr := head

	for !mp[curr] {
		if curr.Next.Val == val && curr.Val != val {
			n := curr.Next
			curr.Next = &Node{
				Val:  x,
				Next: n,
			}
			return head
		}
		if mp[curr.Next] {
			n := curr.Next
			curr.Next = &Node{
				Val:  x,
				Next: n,
			}
			return head
		}

		mp[curr] = true
		curr = curr.Next
	}

	return head
}

func getValues(head *Node) (int, int) {
	var (
		maxVal int
		minVal = math.MaxInt32
		mp     = make(map[*Node]bool)
	)

	for !mp[head] {
		maxVal = max(maxVal, head.Val)
		minVal = min(minVal, head.Val)
		mp[head] = true
		head = head.Next
	}

	return maxVal, minVal
}
