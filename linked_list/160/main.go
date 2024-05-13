package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(head1, head2 *ListNode) *ListNode {
	var (
		l1, l2       int
		curr1, curr2 = head1, head2
	)

	for curr1.Next != nil {
		l1++
		curr1 = curr1.Next
	}

	for curr2.Next != nil {
		l2++
		curr2 = curr2.Next
	}

	if curr1 != curr2 {
		return nil
	}

	curr2 = head2
	for l1 < l2 {
		curr2 = curr2.Next
		l2--
	}

	curr1 = head1
	for l1 > l2 {
		curr1 = curr1.Next
		l1--
	}

	for curr1 != curr2 {
		curr1 = curr1.Next
		curr2 = curr2.Next
	}

	return curr1
}
