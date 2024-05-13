package main

import "fmt"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	fast, slow := head, head

	var prev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}

	prev.Next = nil
	r := mergeSort(slow)
	l := mergeSort(head)

	// fmt.Println(head, slow)
	// fmt.Println(l, r, "suka")
	val := merge(l, r)
	// Print(val)
	return val
}

func Print(head *ListNode) {

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

	fmt.Println()
}

func merge(head1, head2 *ListNode) *ListNode {
	// fmt.Println(head1.Val, head2.Val, "in merge")

	result := &ListNode{}
	tmp := result

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			tmp.Next = head1
			tmp = tmp.Next
			head1 = head1.Next
		} else {
			tmp.Next = head2
			tmp = tmp.Next
			head2 = head2.Next
		}
	}

	if head1 != nil {
		tmp.Next = head1
	}

	if head2 != nil {
		tmp.Next = head2
	}

	return result.Next
}
