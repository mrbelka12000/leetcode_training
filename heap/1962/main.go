package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(minStoneSum([]int{5, 4, 9}, 2))

}

func minStoneSum(piles []int, k int) int {
	pq := PQ(piles)

	heap.Init(&pq)
	var sum int

	for i := 0; i < len(piles); i++ {
		sum += piles[i]
	}

	for i := 0; i < k; i++ {
		val := heap.Pop(&pq).(int)
		x := math.Ceil(float64(val) / 2.0)
		heap.Push(&pq, int(x))
		sum -= val - int(x)
	}

	return sum
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
