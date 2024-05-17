package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	var size int

	curr := head
	for curr != nil {
		size++
		curr = curr.Next
	}

	fill := size/k > 0
	var mod int

	if fill {
		mod = size % k
		size /= k
	}

	curr = head

	result := make([]*ListNode, k)
	for i := 0; i < k; i++ {
		result[i] = &ListNode{}
	}
	tmp := make([]*ListNode, len(result))

	copy(tmp, result)

	var ind int
	for curr != nil {
		if fill {
			for i := 0; i < size; i++ {
				tmp[ind].Next = &ListNode{
					Val: curr.Val,
				}
				tmp[ind] = tmp[ind].Next
				curr = curr.Next
				// fmt.Println(tmp[ind], ind)
			}
			if mod != 0 {
				tmp[ind].Next = &ListNode{
					Val: curr.Val,
				}
				tmp[ind] = tmp[ind].Next
				curr = curr.Next
				mod--
			}
		} else {
			tmp[ind].Next = &ListNode{
				Val: curr.Val,
			}
			tmp[ind] = tmp[ind].Next
			curr = curr.Next
		}

		ind++
	}

	for i := 0; i < k; i++ {
		result[i] = result[i].Next
	}

	return result
}
