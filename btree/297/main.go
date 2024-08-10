package main

import (
	"fmt"
	"strconv"
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

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	path := preOrder(root)

	return path
}

func preOrder(root *TreeNode) (result string) {
	if root == nil {
		return "null"
	}
	result += fmt.Sprint(root.Val) + ","
	result += preOrder(root.Left) + ","
	result += preOrder(root.Right)

	return result
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	path = strings.Split(data, ",")

	return runner()
}

var path []string

func runner() *TreeNode {

	next := path[0]
	path = path[1:]

	if next == "null" {
		return nil
	}
	val, _ := strconv.Atoi(next)
	root := &TreeNode{
		Val: val,
	}

	root.Left = runner()
	root.Right = runner()

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
