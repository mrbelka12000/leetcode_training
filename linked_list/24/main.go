package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	curr := head
	var last, final *ListNode

	for curr != nil {
		prev := reverse(curr, last, 2)
		if prev == nil {
			break
		}
		if final == nil {
			final = prev
		}
		last = curr
		curr = curr.Next
	}

	return final
}

func reverse(head, last *ListNode, k int) *ListNode {
	curr := head
	checker := head
	var prev, tmp *ListNode
	var ind int
	for checker != nil {
		ind++
		checker = checker.Next
	}

	for curr != nil {
		k--
		next := curr.Next
		tmp = next
		curr.Next = prev
		prev = curr
		curr = next
		if k == 0 {
			break
		}
	}

	if last != nil {
		last.Next = prev
	}
	// fmt.Printf("head=%v, start of new list=%v, last=%v, tmp=%v\n", head, prev, last, tmp)

	head.Next = tmp

	return prev
}
