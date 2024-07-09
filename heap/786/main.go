package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(kthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3))
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	pq := &PQ{}
	heap.Init(pq)

	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			heap.Push(pq, []int{arr[i], arr[j]})

			if pq.Len() > k {
				heap.Pop(pq)
			}
		}
	}

	return heap.Pop(pq).([]int)
}

type PQ [][]int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	return float64(pq[i][0])/float64(pq[i][1]) > float64(pq[j][0])/float64(pq[j][1])
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
