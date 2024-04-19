package main

import "container/heap"

func main() {

}

func findClosestElements(arr []int, k int, x int) []int {
	pq := PQ(arr)

	X = x
	heap.Init(&pq)
	var result []int
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(&pq).(int))
	}

	return result
}

var X int

type PQ []int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// max-heap
func (pq PQ) Less(i, j int) bool {
	return (abs(pq[i]-X) < abs(pq[j]-X)) || ((abs(pq[i]-X) == abs(pq[j]-X)) && pq[i] < pq[j])
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

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
