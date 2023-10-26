package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {

	var check bool
	var arr []int
	var num int
	for head != nil {
		if head.Val == 0 {
			if check {
				arr = append(arr, num)
				num = 0
			} else {
				check = true
			}
		}
		if check {
			num += head.Val
		}
		head = head.Next
	}

	result := &ListNode{}
	tmp := result

	for _, v := range arr {
		result.Next = &ListNode{
			Val: v,
		}

		result = result.Next
	}

	return tmp.Next
}
