package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, s int, d int) string {
	l := lca(root, s, d)

	result = strings.Builder{}
	path := make([]byte, 0, 10000)

	getToS(l, s, path)
	getToD(l, d, path)

	return result.String()
}

var result strings.Builder

func getToD(root *TreeNode, d int, path []byte) bool {
	if root == nil {
		return false
	}

	if root.Val == d {
		result.Write(path)
		return true
	}

	if getToD(root.Left, d, append(path, 'L')) {
		return true
	}

	if getToD(root.Right, d, append(path, 'R')) {
		return true
	}

	return false
}

func getToS(root *TreeNode, s int, path []byte) bool {
	if root == nil {
		return false
	}

	if root.Val == s {
		result.Write(path)
		return true
	}

	if getToS(root.Left, s, append(path, 'U')) {
		return true
	}

	if getToS(root.Right, s, append(path, 'U')) {
		return true
	}

	return false
}

func lca(root *TreeNode, s, d int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == s || root.Val == d {
		return root
	}

	l := lca(root.Left, s, d)
	r := lca(root.Right, s, d)
	if l != nil && r != nil {
		return root
	}

	if l != nil {
		return l
	}

	return r
}
