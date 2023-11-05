package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	return bfs(root)
}

func bfs(root *TreeNode) int {
	var result int

	q := []*TreeNode{root}

	for len(q) != 0 {
		var check int
		size := len(q)

		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}

			check += n.Val
		}
		result = check
	}

	return result
}
