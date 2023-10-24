package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if root.Val < val {
		if root.Right == nil {
			root.Right = &TreeNode{
				Val: val,
			}
			return root
		}
		insertIntoBST(root.Right, val)
	}
	if root.Val > val {
		if root.Left == nil {
			root.Left = &TreeNode{
				Val: val,
			}
			return root
		}
		insertIntoBST(root.Left, val)
	}

	return root
}
