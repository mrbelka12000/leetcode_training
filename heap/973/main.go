package main

import (
	"container/heap"
)

func main() {

}

func kClosest(points [][]int, k int) [][]int {
	pq := &PQ{}

	heap.Init(pq)

	for i := 0; i < len(points); i++ {
		heap.Push(pq, points[i])
	}

	var result [][]int
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(pq).([]int))
	}

	return result
}

func getDist(c1 []int) int {
	return c1[0]*c1[0] + c1[1]*c1[1]
}

type PQ [][]int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// min-heap
func (pq PQ) Less(i, j int) bool {
	return getDist(pq[i]) < getDist(pq[j])
}

func (pq *PQ) Push(x interface{}) {
	tmp := x.([]int)
	*pq = append(*pq, tmp)
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	tmp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return tmp
}
