package main

import "container/heap"

func main() {

}

type KthLargest struct {
	pq *PQ
	k  int
}

func Constructor(k int, nums []int) KthLargest {
	pq := &PQ{}

	heap.Init(pq)
	s := KthLargest{
		pq: pq,
		k:  k,
	}

	for i := 0; i < len(nums); i++ {
		s.Add(nums[i])
	}
	return s
}

func (s *KthLargest) Add(num int) int {
	heap.Push(s.pq, num)

	if s.pq.Len() > s.k {
		heap.Pop(s.pq)
	}

	return (*s.pq)[0]
}

type PQ []int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

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
