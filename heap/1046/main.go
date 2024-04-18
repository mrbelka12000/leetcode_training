package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))

}

func lastStoneWeight(stones []int) int {
	pq := &PQ{}
	heap.Init(pq)

	for i := 0; i < len(stones); i++ {
		heap.Push(pq, stones[i])
	}

	for pq.Len() != 1 {
		val1, val2 := heap.Pop(pq).(int), heap.Pop(pq).(int)

		heap.Push(pq, val1-val2)
	}
	return (*pq)[0]
}

type PQ []int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// max-heap
func (pq PQ) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq *PQ) Push(x interface{}) {
	tmp := x.(int)
	*pq = append(*pq, tmp)
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	tmp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return tmp
}
