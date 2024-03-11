package main

import "fmt"

func MyCanFinish(numCourses int, pr [][]int) bool {
	if len(pr) == 0 {
		return true
	}

	g := newGraph()

	for _, v := range pr {
		g.addAdj(v[1], v[0])
	}

	if !g.isValid() {
		return false
	}

	// g.Print()

	for i := 0; i < len(g.store); i++ {
		if g.store[i].cons == 0 {
			continue
		}
		if runner(g.store[i], make(map[int]struct{})) {
			i = -1
		}
	}

	var result int

	for _, v := range g.store {
		if v.cons == 0 {
			result++
		}
	}

	// fmt.Println(result, len(g.store))

	return result == len(g.store)
}

func runner(n *node, visited map[int]struct{}) bool {
	visited[n.val] = struct{}{}
	if n.cons == 0 {
		return true
	}

	for _, v := range n.adj {
		if _, ok := visited[v.val]; ok {
			continue
		}

		if runner(v, visited) {
			n.cons--
			delete(n.adj, v.val)
		}
	}

	return n.cons == 0
}

type (
	node struct {
		val     int
		cons    int
		adj     map[int]*node
		visited bool
	}
	graph struct {
		store    []*node
		finished map[int]*node
	}
)

func newGraph() *graph {
	return &graph{
		finished: make(map[int]*node),
	}
}

func (g *graph) addNode(val int) *node {
	n := &node{
		val: val,
		adj: make(map[int]*node),
	}

	g.store = append(g.store, n)

	return n
}

func (g *graph) getNode(val int) *node {
	for _, v := range g.store {
		if v.val == val {
			return v
		}
	}

	return g.addNode(val)
}

func (g *graph) addAdj(from, to int) {
	fromNode, toNode := g.getNode(from), g.getNode(to)

	fromNode.cons++

	fromNode.adj[to] = toNode
}

func (g *graph) Print() {
	for _, node := range g.store {
		fmt.Printf("node=%v cons=%v has adjs\n", node.val, node.cons)
		for _, adj := range node.adj {
			fmt.Printf("%+v\n", adj.val)
		}
	}
}

func (g *graph) isValid() bool {
	for _, v := range g.store {
		if v.cons == 0 {
			return true
		}
	}

	return false
}
