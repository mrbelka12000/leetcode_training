package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {

	if head.Next == nil {
		return
	}

	fast, slow := head, head
	var last *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		last = slow
		slow = slow.Next
	}

	if last != nil {
		last.Next = nil
	}

	reversed := reverse(slow)
	slow = head

	tmp := &ListNode{}
	// result := &ListNode{
	//     Next: tmp,
	// }

	for slow != nil {
		tmp.Next = slow
		tmp = tmp.Next
		slow = slow.Next

		tmp.Next = reversed
		tmp = tmp.Next
		if reversed != nil {
			reversed = reversed.Next
		}
	}

	return
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
