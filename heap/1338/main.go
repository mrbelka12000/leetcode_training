package main

import "container/heap"

func main() {

}

func minSetSize(arr []int) int {
	mp := make(map[int]int)
	for _, v := range arr {
		mp[v]++
	}

	pq := &PQ{}
	heap.Init(pq)

	for k, v := range mp {
		heap.Push(pq, Info{
			val: k,
			cnt: v,
		})
	}

	var removed, result int

	for pq.Len() >= 0 && removed < len(arr)/2 {
		val := heap.Pop(pq).(Info)
		removed += val.cnt
		result++
	}

	return result
}

type Info struct {
	val int
	cnt int
}

type PQ []Info

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// max-heap
func (pq PQ) Less(i, j int) bool {
	return pq[i].cnt > pq[j].cnt
}

func (pq *PQ) Push(x interface{}) {
	tmp := x.(Info)
	*pq = append(*pq, tmp)
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	tmp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return tmp
}
