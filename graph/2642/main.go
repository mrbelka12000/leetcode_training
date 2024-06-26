package main

import (
	"fmt"
	"math"
)

func main() {
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

type Graph struct {
	store map[int]map[int]int
	n     int
}

func Constructor(n int, edges [][]int) Graph {
	store := make(map[int]map[int]int)
	for _, v := range edges {
		if store[v[0]] == nil {
			store[v[0]] = make(map[int]int)
		}
		store[v[0]][v[1]] = v[2]
	}

	return Graph{
		store: store,
		n:     n,
	}
}

func (this *Graph) AddEdge(edge []int) {
	if this.store[edge[0]] == nil {
		this.store[edge[0]] = make(map[int]int)
	}

	this.store[edge[0]][edge[1]] = edge[2]
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	calc := this.getCalc()
	//p := PQ{}
	//heap.Init(&p)
	calc[node1] = 0
	// p.Push([2]int{node1, 0})

	//heap.Push(&p, [2]int{node1, 0})
	q := [][]int{{node1, 0}}
	for len(q) > 0 {
		item := q[0]
		q = q[1:]

		if item[1] > calc[item[0]] {
			continue
		}

		if item[0] == node2 {
			break
		}

		for k, v := range this.store[item[0]] {
			if calc[k] > calc[item[0]]+v {
				calc[k] = calc[item[0]] + v
				// p.Push([2]int{k, calc[k]})
				//heap.Push(&p, [2]int{k, calc[k]})
				q = append(q, []int{k, calc[k]})
			}
		}
	}

	if calc[node2] == math.MaxInt32 {
		return -1
	}
	return calc[node2]
}

func (this *Graph) getCalc() map[int]int {
	mp := make(map[int]int)

	for i := 0; i < this.n; i++ {
		mp[i] = math.MaxInt32
	}

	return mp
}

func (this *Graph) Print() {
	for k, v := range this.store {
		fmt.Printf("node %v has edjes:\n", k)
		for key, val := range v {
			fmt.Printf("adj=%v has cost=%v\n", key, val)
		}
		fmt.Println()
	}
}

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */
