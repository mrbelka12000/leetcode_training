package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	mp := make(map[int][]int)
	var (
		minVal = math.MaxInt32
		maxVal int
	)

	q := []Info{Info{
		root, 0,
	}}

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]
			node := n.node
			mp[n.col] = append(mp[n.col], node.Val)
			minVal = min(minVal, n.col)
			maxVal = max(maxVal, n.col)
			if node.Left != nil {
				q = append(q, Info{node.Left, n.col - 1})
			}
			if node.Right != nil {
				q = append(q, Info{node.Right, n.col + 1})
			}
		}
	}

	var result [][]int

	for i := minVal; i <= maxVal; i++ {
		result = append(result, mp[i])
	}

	return result
}

type Info struct {
	node *TreeNode
	col  int
}
