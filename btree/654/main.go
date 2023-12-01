package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	root := &TreeNode{}

	dfs(root, nums)

	return root
}

func dfs(root *TreeNode, nums []int) {
	if root == nil || len(nums) == 0 {
		return
	}

	max, ind := maxArr(nums)
	if ind == -1 {
		return
	}
	root.Val = max

	root.Left = crt(nums[:ind])
	root.Right = crt(nums[ind+1:])

	dfs(root.Left, nums[:ind])
	dfs(root.Right, nums[ind+1:])
}

func crt(nums []int) *TreeNode {

	max, ind := maxArr(nums)
	if ind == -1 {
		return nil
	}

	return &TreeNode{
		Val: max,
	}
}

func maxArr(arr []int) (max int, ind int) {
	ind = -1
	max = -1
	for i, v := range arr {
		if v > max {
			max = v
			ind = i
		}
	}

	return
}
