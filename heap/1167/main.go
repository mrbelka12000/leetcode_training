package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(connectSticks([]int{1, 8, 3, 5}))
}

func connectSticks(sticks []int) int {
	pq := PQ(sticks)
	heap.Init(&pq)

	var result int
	for pq.Len() != 1 {
		val1 := heap.Pop(&pq).(int)
		val2 := heap.Pop(&pq).(int)
		heap.Push(&pq, val1+val2)
		result += val1 + val2
	}

	return result
}

type PQ []int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// min-heap
func (pq PQ) Less(i, j int) bool {
	return pq[i] < pq[j]
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
