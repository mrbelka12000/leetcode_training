package main

import (
	"fmt"
	"math"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	dep, diam := getInfo(root)
	res := make([][]string, dep)

	for i := 0; i < len(res); i++ {
		res[i] = make([]string, diam)
	}

	start := (diam - 1) / 2

	res[0][start] = fmt.Sprint(root.Val)

	printer(root, res, 0, start)

	return res
}

/*
[["","","","1","","",""]
,["","","2","","3","",""]
,["","","","4","","",""]]
*/
func printer(root *TreeNode, res [][]string, x, y int) {
	if root == nil {
		return
	}

	if x != 0 {
		res[x][y] = fmt.Sprint(root.Val)
	}

	pos := int(math.Pow(2, float64(len(res)-1-x-1)))

	printer(root.Left, res, x+1, y-pos)
	printer(root.Right, res, x+1, y+pos)
}

func getInfo(root *TreeNode) (dep int, diam int) {
	q := []*TreeNode{root}

	for len(q) != 0 {
		dep++
		size := len(q)
		for i := 0; i < size; i++ {

			n := q[0]
			q = q[1:]
			if n == nil {
				continue
			}

			q = append(q, n.Left)
			q = append(q, n.Right)
		}
	}

	return dep - 1, int(math.Pow(2, float64(dep-1))) - 1
}
