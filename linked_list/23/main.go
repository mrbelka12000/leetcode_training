package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)

	if n == 0 {
		return nil
	}

	if n == 1 {
		return lists[0]
	}

	firstHalf := mergeKLists(lists[:n/2])
	secondHalf := mergeKLists(lists[n/2:])

	return sort(firstHalf, secondHalf)
}

func sort(l, r *ListNode) *ListNode {
	if l == nil && r == nil {
		return nil
	}

	if r == nil {
		return l
	}

	if l == nil {
		return r
	}

	if l.Val <= r.Val {
		prev := l
		for l.Next != nil && l.Next.Val < r.Val {
			l = l.Next
		}
		lNext := l.Next
		l.Next = r
		r.Next = sort(lNext, r.Next)
		return prev
	}

	prev := r
	for r.Next != nil && r.Next.Val < l.Val {
		r = r.Next
	}
	rNext := r.Next
	r.Next = l
	l.Next = sort(l.Next, rNext)
	return prev
}
