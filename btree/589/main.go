package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var runner func(node *Node) []int

	runner = func(node *Node) []int {
		if node == nil {
			return nil
		}

		arr := []int{node.Val}
		for _, v := range node.Children {
			arr = append(arr, runner(v)...)
		}

		return arr
	}

	return runner(root)
}
