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
		s = &ListNode{}
	)

	for head != nil && head.Next != nil {
		s.Next = head
		s = s.Next
		head.Next = head.Next.Next
		head = head.Next
	}

	return result.Next
}
