package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
	fmt.Println(allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}))
}

type n struct {
	id      int
	adj     []*n
	visited bool
	root    *n
}

type g struct {
	nodes []*n
	mp    map[int]*n
}

var paths [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	paths = nil

	g := build()
	for i := 0; i < len(graph); i++ {
		g.addNode(i)
	}
	for i := 0; i < len(graph); i++ {
		for _, v := range graph[i] {
			g.addAdj(i, v)
		}
	}

	path := []int{g.nodes[0].id}
	g.dfs(g.nodes[0], path)

	return paths
}

func (g *g) dfs(n *n, path []int) bool {

	if n.id == g.nodes[len(g.nodes)-1].id {
		getPath(n)
		return true
	}

	for i := 0; i < len(n.adj); i++ {
		n.adj[i].root = n
		g.dfs(n.adj[i], path)
	}

	return false
}

func getPath(n *n) {
	var path []int
	for n != nil {
		path = append(path, n.id)
		n = n.root
	}
	var (
		left  int
		right = len(path) - 1
	)

	for left < right {
		path[left], path[right] = path[right], path[left]
		left++
		right--
	}

	paths = append(paths, path)
}

func build() *g {
	return &g{
		mp: make(map[int]*n),
	}
}

func (g *g) getNode(id int) *n {
	val, ok := g.mp[id]
	if !ok {
		return nil
	}
	return val
}

func (g *g) addNode(id int) {
	val := g.getNode(id)
	if val != nil {
		return
	}
	node := &n{id: id}
	g.nodes = append(g.nodes, node)
	g.mp[id] = node
}

func (g *g) addAdj(from, to int) {
	fromNode := g.getNode(from)
	if fromNode == nil {
		log.Fatal("no from node")
	}

	toNode := g.getNode(to)
	if toNode == nil {
		log.Fatal("no to node")
	}

	fromNode.adj = append(fromNode.adj, toNode)
}
