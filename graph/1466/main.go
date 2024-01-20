package main

import "fmt"

func main() {
	fmt.Println(minReorder(6, [][]int{
		{0, 1},
		{1, 3},
		{2, 3},
		{4, 0},
		{4, 5}}))
}

func minReorder(n int, connections [][]int) int {
	g := NewGraph()

	for _, v := range connections {
		g.AddEdge(v[0], v[1])
	}

	return bfs(g.nodes[0])
}

func bfs(start *Node) int {
	var count int

	q := []*Node{start}
	visited := make(map[int]struct{})
	for len(q) != 0 {

		node := q[0]
		q = q[1:]

		visited[node.Val] = struct{}{}

		for k, v := range node.FakeEdges {
			if _, ok := visited[k]; ok {
				continue
			}

			if _, ok := v.Edges[node.Val]; !ok {
				// fmt.Println(node.Val, k)
				count++
			}

			q = append(q, v)
		}
	}

	return count
}

type (
	Graph struct {
		nodes map[int]*Node
	}

	Node struct {
		Val       int
		Edges     map[int]*Node
		FakeEdges map[int]*Node
	}
)

func NewGraph() *Graph {
	return &Graph{
		nodes: map[int]*Node{},
	}
}

func (g *Graph) addNode(val int) *Node {
	node, ok := g.nodes[val]
	if !ok {
		node = &Node{
			Val:       val,
			Edges:     map[int]*Node{},
			FakeEdges: map[int]*Node{},
		}
		g.nodes[val] = node

		return node
	}
	return node
}

func (g *Graph) AddEdge(from, to int) {
	fromNode, toNode := g.nodes[from], g.nodes[to]
	if fromNode == nil {
		fromNode = g.addNode(from)
	}
	if toNode == nil {
		toNode = g.addNode(to)
	}

	fromNode.Edges[to] = toNode

	fromNode.FakeEdges[to] = toNode
	toNode.FakeEdges[from] = fromNode
}
