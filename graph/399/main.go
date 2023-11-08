package main

import (
	"fmt"
)

func main() {
	fmt.Println(calcEquation(
		[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
		[]float64{1.5, 2.5, 5.0},
		[][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}),
	)
}

func calcEquation(eq [][]string, values []float64, que [][]string) []float64 {
	g := &graph{}
	result := make([]float64, len(que))

	// init graph
	for i := 0; i < len(eq); i++ {
		g.addAdj(eq[i][0], eq[i][1], values[i])
		g.addAdj(eq[i][1], eq[i][0], 1/values[i])
	}

	for i := 0; i < len(que); i++ {
		result[i] = g.math(que[i][0], que[i][1])
		fmt.Println()
	}

	//g.print()

	return result
}

type (
	graph struct {
		nodes []*node
	}

	node struct {
		Key string
		Val string
		Adj []*adj
	}
	adj struct {
		n    *node
		cost float64
	}
)

// addNode
func (g *graph) addNode(key string) *node {
	if g.getNode(key) != nil {
		return nil
	}

	n := &node{
		Key: key,
	}
	g.nodes = append(g.nodes, n)
	return n
}

// getNode
func (g *graph) getNode(key string) *node {
	for _, v := range g.nodes {
		if v.Key == key {
			return v
		}
	}
	return nil
}

// addAdj
func (g *graph) addAdj(from, to string, cost float64) {
	fromNode, toNode := g.getNode(from), g.getNode(to)
	if fromNode == nil {
		fromNode = g.addNode(from)
	}
	if toNode == nil {
		toNode = g.addNode(to)
	}

	fromNode.Adj = append(fromNode.Adj, &adj{
		n:    toNode,
		cost: cost,
	})
}

// print
func (g *graph) print() {

	for _, v := range g.nodes {
		fmt.Printf("Node: %+v has adj: \n", v)
		for _, a := range v.Adj {
			fmt.Printf("Adj: %+v\n", a)
		}
		fmt.Println()
	}
}

// math
func (g *graph) math(from, to string) float64 {
	fromNode, toNode := g.getNode(from), g.getNode(to)
	if fromNode == nil || toNode == nil {
		return -1
	}

	mp := make(map[string]struct{})

	val, ok := dfs(fromNode, toNode, 1, mp)
	if !ok {
		return -1
	}
	//fmt.Println(val)
	return val
}

func dfs(from, to *node, path float64, mp map[string]struct{}) (float64, bool) {
	if _, ok := mp[from.Key]; ok {
		return 0, false
	}

	if from.Key == to.Key {
		return path, true
	}

	//fmt.Println(from.Key)
	mp[from.Key] = struct{}{}

	for _, v := range from.Adj {
		result, ok := dfs(v.n, to, path*v.cost, mp)
		if ok {
			return result, true
		}
	}

	return 0, false
}
