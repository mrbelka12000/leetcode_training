package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	runner(root)
}

func runner(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := runner(root.Left)
	right := runner(root.Right)

	// fmt.Println(left, right, root)

	rightMost := getRight(left)
	// fmt.Println(rightMost, "suka")
	if rightMost != nil {
		rightMost.Right = right
	}

	if left != nil {
		root.Right = left
	}

	root.Left = nil

	return root
}

func getRight(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Right == nil {
		return root
	}

	return getRight(root.Right)
}
