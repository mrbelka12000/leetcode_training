package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	totalSum := getTotalSum(root)
	result = 0

	dfs(root, totalSum)

	return result % MOD
}

const MOD int = 1000000007

func dfs(root *TreeNode, totalSum int) int {
	if root == nil {
		return 0
	}

	sum := dfs(root.Left, totalSum) + dfs(root.Right, totalSum) + root.Val
	result = max(((totalSum - sum) * sum), result)
	return sum
}

var result int

func getTotalSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return getTotalSum(root.Left) + getTotalSum(root.Right) + root.Val
}
