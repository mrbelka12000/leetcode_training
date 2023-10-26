package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var (
		result = &ListNode{
			Next: head,
		}
		next = head.Next
	)

	for head != nil && next != nil {

		break
		head = head.Next
		next = next.Next
	}

	return result.Next
}
