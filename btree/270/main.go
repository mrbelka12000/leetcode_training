package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {

	var (
		minDist = float64(math.MaxInt16)
		result  = math.MaxInt32
		runner  func(root *TreeNode)
	)

	runner = func(curr *TreeNode) {
		if curr == nil {
			return
		}

		if minDist == abs(target-float64(curr.Val)) {
			result = min(result, curr.Val)
		}
		if minDist > abs(target-float64(curr.Val)) {
			result = curr.Val
			minDist = abs(target - float64(curr.Val))
		}

		if float64(curr.Val) >= target {
			runner(curr.Left)
		} else {
			runner(curr.Right)
		}
	}

	runner(root)

	return result
}

func abs(a float64) float64 {
	if a < 0 {
		return a * -1
	}
	return a
}
