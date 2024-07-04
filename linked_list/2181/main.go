package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result
	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Val == 0 {
			tmp.Next = &ListNode{}
			tmp = tmp.Next
		} else {
			tmp.Val += curr.Val
		}
		curr = curr.Next
	}

	return result.Next
}
