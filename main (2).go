package main

import (
	"fmt"
	"math"
)

func main() {

	// construct graph
	graph := map[int][]Edge{
		1: {
			{
				To:   2,
				Cost: 5,
			}, {
				To:   3,
				Cost: 2,
			},
		},
		2: {{
			To:   1,
			Cost: 5,
		}, {
			To:   3,
			Cost: 8,
		}, {
			To:   4,
			Cost: 4,
		}, {
			To:   5,
			Cost: 2,
		},
		},
		3: {
			{
				To:   1,
				Cost: 2,
			}, {
				To:   2,
				Cost: 8,
			}, {
				To:   5,
				Cost: 7,
			},
		},
		4: {
			{
				To:   2,
				Cost: 4,
			},
			{
				To:   5,
				Cost: 6,
			},
			{
				To:   6,
				Cost: 3,
			},
		},
		5: {
			{
				To:   2,
				Cost: 2,
			},
			{
				To:   3,
				Cost: 7,
			},
			{
				To:   4,
				Cost: 6,
			},
			{
				To:   6,
				Cost: 1,
			},
		},
		6: {
			{
				To:   4,
				Cost: 3,
			},
			{
				To:   5,
				Cost: 1,
			},
		},
	}

	fmt.Println(Dijkstra(graph, 1, 6))
}

// Edge information about connections
type Edge struct {
	To   int
	Cost int
}

// Dijkstra algorithm implementation
func Dijkstra(graph map[int][]Edge, start, end int) int {

	// initialize maximum costs to every node
	costs := make(map[int]int)
	for i := 1; i <= len(graph); i++ {
		costs[i] = math.MaxInt32
	}

	// set cost from start to zero, because we start here, so we do not need to calculate cost to start
	costs[start] = 0

	// initialize visited map for not to cycle
	visited := make(map[int]bool)

	q := []int{start}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]

		// if we get to the end stop runner
		if n == end {
			break
		}

		if visited[n] {
			continue
		}

		// set to visited current node
		visited[n] = true

		// traverse to all connections from current node
		for _, edge := range graph[n] {

			// calculate min cost
			if costs[edge.To] > costs[n]+edge.Cost {
				costs[edge.To] = costs[n] + edge.Cost
				q = append(q, edge.To)
			}
		}
	}

	return costs[end]
}
