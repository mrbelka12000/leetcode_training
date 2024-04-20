package main

import "container/heap"

func main() {

}

func numberGame(nums []int) []int {
	pq := PQ(nums)

	heap.Init(&pq)

	var arr []int

	for pq.Len() > 0 {
		v1 := heap.Pop(&pq).(int)
		v2 := heap.Pop(&pq).(int)

		arr = append(arr, v2)
		arr = append(arr, v1)
	}

	return arr
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
