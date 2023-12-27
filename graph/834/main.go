package main

import "fmt"

func main() {

	fmt.Println(sumOfDistancesInTree(6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}))
}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	g := &G{
		store: make(map[int]*node),
	}

	for _, v := range edges {
		g.AddEdge(v[0], v[1])
		g.AddEdge(v[1], v[0])
	}
	if len(g.store) == 0 {
		return []int{0}
	}

	result := g.CalculateDistance()

	return result
}

type (
	G struct {
		store map[int]*node
	}
	node struct {
		val   int
		edjes []*node
	}
)

func (g *G) AddNode(val int) *node {
	n := &node{val: val}

	g.store[val] = n

	return n
}

func (g *G) GetNode(val int) *node {
	return g.store[val]
}

func (g *G) AddEdge(from, to int) {
	fromNode, toNode := g.GetNode(from), g.GetNode(to)
	if fromNode == nil {
		fromNode = g.AddNode(from)
	}
	if toNode == nil {
		toNode = g.AddNode(to)
	}

	fromNode.edjes = append(fromNode.edjes, toNode)
}

func (g *G) CalculateDistance() []int {
	size := make([]int, len(g.store))
	dp := make([]int, len(g.store))
	result := make([]int, len(g.store))

	dfs(g.GetNode(0), nil, size, dp)
	dfs1(g.GetNode(0), nil, result, dp, size)
	return result
}

func dfs(node, par *node, size, dp []int) {
	size[node.val] = 1
	dp[node.val] = 0

	for _, v := range node.edjes {
		if v != par {
			dfs(v, node, size, dp)
			size[node.val] += size[v.val]
			dp[node.val] += dp[v.val] + size[v.val]
		}
	}
}

func reroot(from, to int, dp, size []int) {
	dp[from] -= size[to] + dp[to]
	size[from] -= size[to]

	size[to] += size[from]
	dp[to] += size[from] + dp[from]
}

func dfs1(node, par *node, ans, dp, size []int) {
	ans[node.val] = dp[node.val]

	for _, v := range node.edjes {
		if par != v {
			reroot(node.val, v.val, dp, size)
			dfs1(v, node, ans, dp, size)
			reroot(v.val, node.val, dp, size)
		}
	}
}
