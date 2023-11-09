package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int

	collectPaths(root, &result, []int{}, targetSum, root.Val)

	return result
}

func collectPaths(root *TreeNode, result *[][]int, arr []int, targetSum, currSum int) {
	if root == nil {
		return
	}
	// currSum += root.Val

	if root.Left == nil && root.Right == nil {
		if targetSum == currSum {
			arr = append(arr, root.Val)
			*result = append(*result, arr)
		}
		return
	}

	check := make([]int, len(arr))
	if root.Left != nil {
		copy(check, arr)
		collectPaths(root.Left, result, append(check, root.Val), targetSum, currSum+root.Left.Val)
	}
	if root.Right != nil {
		copy(check, arr)
		collectPaths(root.Right, result, append(check, root.Val), targetSum, currSum+root.Right.Val)
	}
}
