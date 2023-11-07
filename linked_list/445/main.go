package main

import "fmt"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	result := &ListNode{}

	l1, ln1 := getLen(l1)
	l2, ln2 := getLen(l2)

	var prev *ListNode
	var sum, ost int
	var arr []*ListNode

	for l1 != nil && l2 != nil {
		sum = 0

		if ln1 > ln2 {
			if l1 != nil {
				sum += l1.Val
				l1 = l1.Next
			}
			ln1--
		} else if ln1 < ln2 {
			if l2 != nil {
				sum += l2.Val
				l2 = l2.Next
			}
			ln2--
		} else {
			if l1 != nil {
				sum += l1.Val
				l1 = l1.Next
			}
			ln1--
			if l2 != nil {
				sum += l2.Val
				l2 = l2.Next
			}
			ln2--
		}

		// fmt.Println(sum, ost,"before")

		sum += ost
		ost = sum / 10
		sum %= 10

		inc := ost != 0

		// fmt.Println(sum, ost,inc,"after")

		curr := &ListNode{
			Val: sum,
		}

		if inc {
			if prev == nil {
				prev = &ListNode{
					Val: ost,
				}
				result.Next = prev
				result = result.Next
				arr = append(arr, prev)

				ost = 0
			} else {
				prev.Val++
				ost = 0
			}
		}

		result.Next = curr
		arr = append(arr, curr)
		prev = curr
		result = result.Next
		// print(check.Next)
	}

	return incr(arr)
}

func getLen(l *ListNode) (*ListNode, int) {
	var check int
	temp := &ListNode{
		Next: l,
	}
	for l != nil {
		l = l.Next
		check++
	}

	return temp.Next, check
}

func print(l *ListNode) {
	check := &ListNode{
		Next: l,
	}
	for check.Next != nil {
		fmt.Print(check.Next.Val, "-->")
		check = check.Next
	}
	fmt.Println()
}

func incr(arr []*ListNode) *ListNode {
	// fmt.Println(arr[len(arr)-2].Val)
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i].Val == 10 {
			arr[i].Val = 0
			if i == 0 {
				tmp := &ListNode{
					Val:  1,
					Next: arr[0],
				}
				// fmt.Println("popal")
				return tmp
			}
			arr[i-1].Val++
		}
	}
	return arr[0]
}
