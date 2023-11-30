package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var result []int

	return inorder(root, k, &result)
}

func inorder(root *TreeNode, k int, arr *[]int) int {
	if root == nil {
		return -1
	}

	val := inorder(root.Left, k, arr)
	if val != -1 {
		return val
	}
	*arr = append(*arr, root.Val)

	val = inorder(root.Right, k, arr)
	if val != -1 {
		return val
	}

	if len(*arr)-k >= 0 {
		return (*arr)[k-1]
	}

	return -1
}
