package main

import "container/heap"

func main() {

}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	pq := PQ(nums)
	heap.Init(&pq)
	var result int
	prev := heap.Pop(&pq).(int)
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(int)
		if prev == curr {
			continue
		}

		result = max(result, curr-prev)
		prev = curr
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
