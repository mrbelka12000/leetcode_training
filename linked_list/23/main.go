package main

import "sort"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	result := &ListNode{
		Next: &ListNode{},
	}
	var nums []int

	for {

		var check int

		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				check++
				continue
			}
			nums = append(nums, lists[i].Val)
			lists[i] = lists[i].Next
		}

		if check == len(lists) {
			break
		}
	}
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)

	tmp := result

	for _, v := range nums {
		curr := &ListNode{
			Val: v,
		}
		result.Next = curr
		result = result.Next
	}

	return tmp.Next
}
