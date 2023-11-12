package main

import "fmt"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	var path []string

	collect(root, getStrVal(root.Val), &path)

	min := path[0]

	for _, v := range path {
		if v < min {
			min = v
		}
	}
	return min
}

func collect(root *TreeNode, path string, arr *[]string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*arr = append(*arr, path)
		return
	}
	if root.Left != nil {
		collect(root.Left, getStrVal(root.Left.Val)+path, arr)
	}
	if root.Right != nil {
		collect(root.Right, getStrVal(root.Right.Val)+path, arr)
	}
}

func getStrVal(val int) string {
	return fmt.Sprint(string(val + 49 + 48))
}
