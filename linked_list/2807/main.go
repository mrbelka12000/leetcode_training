package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		if prev != nil {
			val := GCD(prev.Val, curr.Val)
			node := &ListNode{
				Next: curr,
				Val:  val,
			}
			prev.Next = node
		}

		prev = curr
		curr = curr.Next
	}

	return head
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
