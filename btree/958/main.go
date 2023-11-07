package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	if isLeaf(root) {
		return true
	}

	return isComplete(root)
}

func isComplete(root *TreeNode) bool {
	q := []*TreeNode{root}

	var (
		count       int
		prevLastArr []*TreeNode
	)

	if isLeaf(root.Left) {
		count++
	}
	if isLeaf(root.Right) {
		count++
	}

	for len(q) != 0 {
		size := len(q)

		count = 0
		prevLastArr = q
		isFull := true
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			if isLeaf(n.Left) {
				count++
			}
			if isLeaf(n.Right) {
				count++
			}

			if n.Left != nil {
				q = append(q, n.Left)
			} else if n.Left == nil {
				isFull = false
			}
			if n.Right != nil {
				q = append(q, n.Right)
			} else {
				isFull = false
			}
		}

		isPreLast := count == size*2

		// fmt.Println(isPreLast, count, size, prevLastArr[0].Val)
		if isPreLast {
			break
		}
		if count == size*2 {
			isPreLast = true
		}
		// fmt.Println(len(q), isFull)
		if len(q) != 0 && !isFull {
			return false
		}

	}

	var lastArr []*TreeNode
	for _, v := range prevLastArr {
		lastArr = append(lastArr, v.Left)
		lastArr = append(lastArr, v.Right)
	}

	return isLeft(lastArr)
}

func isLeaf(n *TreeNode) bool {
	return n == nil || (n.Left == nil && n.Right == nil)
}

func isLeft(arr []*TreeNode) bool {
	var (
		between bool
	)
	for _, v := range arr {
		if v == nil {
			between = true
			continue
		}
		if v != nil && between {
			return false
		}

		between = false
	}

	return true
}
