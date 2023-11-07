package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := collect(root)
	reverse(result)

	return result
}

func reverse(s [][]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func collect(root *TreeNode) [][]int {
	var result [][]int

	q := []*TreeNode{root}

	for len(q) != 0 {

		size := len(q)
		var temp []int
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
			temp = append(temp, n.Val)
		}
		result = append(result, temp)
	}

	return result
}
