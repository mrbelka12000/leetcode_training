package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := make([][]Edge, n)

	for _, v := range flights {
		graph[v[0]] = append(graph[v[0]], Edge{To: v[1], Cost: v[2]})
	}
	dp := make([]int, n)
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	dp[src] = 0
	q := []Edge{{
		To: src,
	}}

	for len(q) > 0 && k > -1 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			nn := node.To
			//fmt.Printf("FROM %+v\n", node)
			//fmt.Println(dp)
			for _, edge := range graph[nn] {
				//fmt.Printf("%+v\n", edge)
				if dp[edge.To] > node.Cost+edge.Cost {
					dp[edge.To] = node.Cost + edge.Cost
					q = append(q, Edge{To: edge.To, Cost: dp[edge.To]})
				}
			}
			//fmt.Println()
		}

		k--
	}
	//fmt.Println(dp)

	if dp[dst] == math.MaxInt32 {
		return -1
	}
	return dp[dst]
}

type Edge struct {
	To   int
	Cost int
}

// 0 - 1
//   \ |
//     2
//     |
//     3
