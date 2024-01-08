package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var result int

	runner(root, low, high, &result)

	return result
}

func runner(root *TreeNode, l, h int, ans *int) {
	if root == nil {
		return
	}
	// fmt.Println(root.Val)

	if root.Val >= l && root.Val <= h {
		*ans += root.Val
	}

	if l <= root.Val {
		runner(root.Left, l, h, ans)
	}

	if h >= root.Val {
		runner(root.Right, l, h, ans)
	}
}
