package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	mp := make(map[int]*Node)

	dfs(node, mp, make(map[int]bool))

	check := make(map[int]map[int]*Node)

	for k := range mp {
		check[k] = make(map[int]*Node)
	}

	n := builder(node, nil, mp, make(map[int]bool), check)

	for k, v := range check {
		for _, n := range v {
			mp[k].Neighbors = append(mp[k].Neighbors, n)
		}
	}

	return n
}

func builder(node, tmp *Node, mp map[int]*Node, used map[int]bool, check map[int]map[int]*Node) *Node {
	tmp = mp[node.Val]
	for _, n := range node.Neighbors {
		if !used[n.Val] {
			used[n.Val] = true
			check[tmp.Val][n.Val] = builder(n, nil, mp, used, check)
		} else {
			check[tmp.Val][n.Val] = mp[n.Val]
		}
	}

	return tmp
}

func dfs(n *Node, mp map[int]*Node, used map[int]bool) {

	used[n.Val] = true
	if mp[n.Val] == nil {
		mp[n.Val] = &Node{Val: n.Val}
	}

	for _, v := range n.Neighbors {
		if !used[v.Val] {
			dfs(v, mp, used)
		}
	}
}
