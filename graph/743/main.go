package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	val := networkDelayTime([][]int{{3, 5, 78}, {2, 1, 1}, {1, 3, 0}, {4, 3, 59}, {5, 3, 85}, {5, 2, 22}, {2, 4, 23}, {1, 4, 43}, {4, 5, 75}, {5, 1, 15}, {1, 5, 91}, {4, 1, 16}, {3, 2, 98}, {3, 4, 22}, {5, 4, 31}, {1, 2, 0}, {2, 5, 4}, {4, 2, 51}, {3, 1, 36}, {2, 3, 59}}, 5, 5)
	if val != 31 {
		fmt.Println(val, "NO")
		return
	}

	fmt.Println("OK")
}

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][]Edge, n+1)

	for _, v := range times {
		graph[v[0]] = append(graph[v[0]], Edge{
			To:   v[1],
			Cost: v[2],
		})
	}

	cost := make(map[int]int)

	for i := 1; i <= n; i++ {
		cost[i] = math.MaxInt32
	}
	cost[k] = 0

	var result int

	pq := &PQ{}
	heap.Init(pq)
	heap.Push(pq, [2]int{k, 0})

	for pq.Len() > 0 {
		item := pq.Pop().([2]int)

		if item[1] > cost[item[0]] {
			continue
		}

		node := item[0]
		for _, edge := range graph[node] {
			if cost[edge.To] > cost[node]+edge.Cost {
				cost[edge.To] = cost[node] + edge.Cost
				heap.Push(pq, [2]int{edge.To, cost[edge.To]})
			}
		}
	}

	for _, v := range cost {
		if v == math.MaxInt32 {
			return -1
		}
		result = max(result, v)
	}

	return result
}

type Edge struct {
	To   int
	Cost int
}

type PQ [][2]int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// min-heap
func (pq PQ) Less(i, j int) bool {
	return pq[i][1] < pq[j][1]
}

func (pq *PQ) Push(x interface{}) {
	tmp := x.([2]int)
	*pq = append(*pq, tmp)
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	tmp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return tmp
}
