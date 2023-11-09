package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	return bfs(root)
}

func bfs(root *TreeNode) [][]int {
	var result [][]int

	q := []*TreeNode{root}

	for len(q) != 0 {
		size := len(q)

		var level []int
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			level = append(level, n.Val)

			if n.Left != nil {
				q = append(q, n.Left)
			}

			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		result = append(result, level)
	}

	return result
}
