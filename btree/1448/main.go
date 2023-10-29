package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	result := 1

	check(root.Left, root.Val, &result)

	check(root.Right, root.Val, &result)

	return result
}

func check(root *TreeNode, max int, count *int) {
	if root == nil {
		return
	}
	if root.Val >= max {
		*count++
		max = root.Val
	}

	check(root.Left, max, count)

	check(root.Right, max, count)
}
