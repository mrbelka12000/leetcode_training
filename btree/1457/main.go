package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	var result int
	arr := [10]int{}
	arr[root.Val]++

	dfs(
		root,
		arr,
		&result,
	)

	return result
}

func dfs(root *TreeNode, arr [10]int, result *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		var check bool
		for _, v := range arr {
			if v%2 != 0 {
				if check {
					return
				}
				check = true
			}
		}
		*result++
		return
	}

	if root.Left != nil {
		arr[root.Left.Val]++
		dfs(root.Left, arr, result)
		arr[root.Left.Val]--
	}

	if root.Right != nil {
		arr[root.Right.Val]++
		dfs(root.Right, arr, result)
		arr[root.Right.Val]--
	}
}
