package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	arr := []int{head.Val}
	curr := head.Next
	var check bool
	for curr != nil {
		if arr[len(arr)-1] == curr.Val {
			check = true
		} else {
			if check {
				arr = arr[:len(arr)-1]
				check = false
			}
			arr = append(arr, curr.Val)
		}
		curr = curr.Next
	}

	if check {
		arr = arr[:len(arr)-1]
	}

	result := &ListNode{
		Next: &ListNode{},
	}

	tmp := result
	tmp.Next.Val = -111

	for _, v := range arr {
		result.Next = &ListNode{
			Val: v,
		}

		result = result.Next
	}

	if tmp.Next.Val == -111 {
		return nil
	}

	return tmp.Next
}
