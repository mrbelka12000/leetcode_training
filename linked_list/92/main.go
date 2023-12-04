package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	ind := 1
	result := &ListNode{
		Next: head,
	}
	curr := head
	var tmp *ListNode
	for curr != nil {
		if ind == left {
			var prev *ListNode
			for ind != right+1 {
				next := curr.Next
				curr.Next = prev
				prev = curr
				curr = next
				ind++
			}

			if tmp != nil {
				tmp.Next = prev
			}

			if left == 1 {
				result.Next = prev
			}

			if tmp == nil && curr == nil {
				return prev
			}
			// return curr

			next := prev
			for next != nil && next.Next != nil {
				next = next.Next
			}
			// fmt.Println(next)
			next.Next = curr

			break
		}

		tmp = curr
		curr = curr.Next
		ind++
	}

	// fmt.Println(result)

	return result.Next
}
