package main

import (
	"container/heap"
)

func main() {

}

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	mp := make(map[int][]info)

	for i, v := range edges {
		mp[v[0]] = append(mp[v[0]], info{
			cost: succProb[i],
			val:  v[1],
		})

		mp[v[1]] = append(mp[v[1]], info{
			cost: succProb[i],
			val:  v[0],
		})

	}

	return shortestPath(mp, start_node, end_node)
}

func shortestPath(graph map[int][]info, start, end int) float64 {
	coster := make(map[int]float64)
	coster[start] = 1

	pq := &PQ{info{val: start, cost: 0}}
	heap.Init(pq)

	for pq.Len() != 0 {

		n := heap.Pop(pq).(info)

		node := n.val
		dist := n.cost
		if dist > coster[node] {
			continue
		}

		if node == end {
			break
		}

		for _, v := range graph[node] {
			newDist := coster[node] * v.cost
			if newDist > coster[v.val] {
				coster[v.val] = newDist
				heap.Push(pq, (info{cost: coster[v.val], val: v.val}))
			}
		}
	}

	return coster[end]
}

type PQ []info

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// min-heap
func (pq PQ) Less(i, j int) bool {
	return pq[i].cost > pq[j].cost
}

func (pq *PQ) Push(x interface{}) {
	tmp := x.(info)
	*pq = append(*pq, tmp)
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	tmp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return tmp
}

type info struct {
	cost float64
	val  int
}
