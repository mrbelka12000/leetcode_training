package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var (
		result [][]int
		q      = []*TreeNode{root}
	)

	dir := true

	for len(q) != 0 {

		size := len(q)
		temp := make([]int, size)
		lastInd := size - 1
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			if dir {
				temp[i] = n.Val
			} else {
				temp[lastInd] = n.Val
				lastInd--
			}

			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}

		dir = dir == false
		result = append(result, temp)
	}

	return result
}
