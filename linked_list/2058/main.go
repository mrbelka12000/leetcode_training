package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	prev := head
	curr := head.Next

	ind := 2
	minDist := math.MaxInt32
	first := -1
	currVal := -1
	for curr != nil && curr.Next != nil {
		next := curr.Next
		if prev.Val < curr.Val && curr.Val > next.Val {
			if first == -1 {
				first = ind
			}
			if currVal != -1 {
				minDist = min(minDist, ind-currVal)
			}
			currVal = ind
		}
		if prev.Val > curr.Val && curr.Val < next.Val {
			if first == -1 {
				first = ind
			}
			if currVal != -1 {
				minDist = min(minDist, ind-currVal)
			}
			currVal = ind
		}
		ind++
		prev = curr
		curr = curr.Next
	}

	if minDist == math.MaxInt32 {
		return []int{-1, -1}
	}

	return []int{minDist, currVal - first}
}
