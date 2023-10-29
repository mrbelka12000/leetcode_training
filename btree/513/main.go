package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var (
		check  bool
		result int
	)

	q := []*TreeNode{root}

	for len(q) != 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			if !check {
				result = q[0].Val
			}
			check = true
			t := q[0]
			q = q[1:]

			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		check = false
	}

	return result
}
