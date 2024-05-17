package main

import (
	"fmt"
	"strings"
)

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	result := &strings.Builder{}

	runner(root, result)

	resStr := result.String()

	return resStr[1 : len(resStr)-1]
}

func runner(root *TreeNode, res *strings.Builder) {

	res.WriteString(fmt.Sprintf("(%v", root.Val))

	if root.Left != nil {
		runner(root.Left, res)
	}

	if root.Right != nil {
		if root.Left == nil {
			res.WriteString("()")
		}
		runner(root.Right, res)
	}

	res.WriteString(")")
}
