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

func findRoot(tree []*Node) *Node {
	var xor int
	for _, v := range tree {
		xor ^= v.Val
		for _, c := range v.Children {
			xor ^= c.Val
		}
	}

	for _, v := range tree {
		if v.Val == xor {
			return v
		}
	}

	return nil
}
