package main

import "sort"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	return bfs(root, k)
}

func bfs(root *TreeNode, k int) int64 {
	q := []*TreeNode{root}

	var sums []int

	for len(q) != 0 {
		size := len(q)

		var sum int

		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}

			sum += n.Val
		}

		sums = append(sums, sum)
	}

	sort.Ints(sums)
	if k > len(sums) {
		return -1
	}

	return int64(sums[len(sums)-k])
}
