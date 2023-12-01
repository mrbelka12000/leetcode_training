package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	var arr []int
	result := &ListNode{
		Next: head,
	}

	curr := head

	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}

	check := head
	ind := 1
	for check != nil {
		if ind == k {
			// fmt.Println(check.Val, arr[len(arr)-k])
			check.Val = arr[len(arr)-k]
		} else if len(arr)-k+1 == ind {
			// fmt.Println(check.Val, arr[k-1])
			check.Val = arr[k-1]
		}
		check = check.Next
		ind++
	}

	return result.Next
}
