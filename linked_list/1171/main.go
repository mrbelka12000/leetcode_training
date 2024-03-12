package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	var arr []int
	curr := head
	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}

	for i := 0; i < len(arr); i++ {
		var tmp int
		for j := i; j < len(arr); j++ {
			tmp += arr[j]
			if tmp == 0 {
				arr = append(arr[:i], arr[j+1:]...)
				i = -1
				break
			}
		}
	}

	tmp := &ListNode{}

	result := tmp

	for _, v := range arr {
		n := &ListNode{
			Val: v,
		}
		tmp.Next = n
		tmp = tmp.Next
	}

	return result.Next
}
