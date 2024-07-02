package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(makePrefSumNonNegative([]int{2, 3, -5, 4}))
}

func makePrefSumNonNegative(nums []int) int {
	var result, ps int
	pq := &PQ{}
	heap.Init(pq)

	for i := 0; i < len(nums); i++ {
		ps += nums[i]
		heap.Push(pq, nums[i])
		if ps < 0 {
			val := heap.Pop(pq).(int)
			ps -= val
			result++
		}
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
