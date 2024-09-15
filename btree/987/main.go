package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello world")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalTraversal(root *TreeNode) [][]int {
	yl, yr, xx := 0, 0, 0
	store := make(map[[2]int][]int)

	var runner func(node *TreeNode, x, y int)

	runner = func(node *TreeNode, x, y int) {
		if node == nil {
			return
		}
		yl = min(yl, y)
		yr = max(yr, y)
		xx = max(xx, x)
		store[[2]int{x, y}] = append(store[[2]int{x, y}], node.Val)

		runner(node.Left, x+1, y-1)
		runner(node.Right, x+1, y+1)
	}

	runner(root, 0, 0)

	var result [][]int

	for y := yl; y <= yr; y++ {
		var tmp []int
		for x := 0; x <= xx; x++ {
			if val, ok := store[[2]int{x, y}]; ok {
				sort.Ints(val)
				tmp = append(tmp, val...)
			}
		}
		result = append(result, tmp)
	}

	return result
}
