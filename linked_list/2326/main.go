package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	result := make([][]int, m)
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, n)
	}

	left, right := 0, n-1
	bot, top := 0, m-1

	for left <= right && bot <= top {
		for i := left; i <= right; i++ {
			val := -1
			if head != nil {
				val = head.Val
				head = head.Next
			}
			result[bot][i] = val
		}

		for i := bot + 1; i <= top; i++ {
			val := -1
			if head != nil {
				val = head.Val
				head = head.Next
			}
			result[i][right] = val
		}

		if bot != top {
			for i := right - 1; i >= left; i-- {
				val := -1
				if head != nil {
					val = head.Val
					head = head.Next
				}
				result[top][i] = val
			}
		}

		if right != left {
			for i := top - 1; i >= bot+1; i-- {
				val := -1
				if head != nil {
					val = head.Val
					head = head.Next
				}
				result[i][left] = val
			}
		}

		left++
		right--
		bot++
		top--
	}

	return result
}
