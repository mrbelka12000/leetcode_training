package main

import (
	"container/heap"
)

func main() {

}

func totalCost(costs []int, k int, cdt int) int64 {
	lPQ := &PQ{}
	rPQ := &PQ{}

	heap.Init(lPQ)
	heap.Init(rPQ)

	var (
		l, r = cdt, len(costs) - cdt - 1
	)
	mp := make(map[int]bool)

	for i := 0; i < cdt; i++ {
		mp[i] = true
		heap.Push(lPQ, costs[i])
	}

	for i := len(costs) - 1; i >= len(costs)-cdt; i-- {
		if mp[i] {
			break
		}
		mp[i] = true

		heap.Push(rPQ, costs[i])
	}
	// fmt.Println(mp)
	var result int64
	for i := 0; i < k; i++ {
		// fmt.Println(lPQ,l)
		// fmt.Println(rPQ,r)
		if lPQ.Len() == 0 {
			result += int64(heap.Pop(rPQ).(int))
			if !mp[r] && r >= 0 {
				heap.Push(rPQ, costs[r])
			}
			mp[r] = true
			r--
			continue
		} else if rPQ.Len() == 0 {
			result += int64(heap.Pop(lPQ).(int))
			if !mp[l] && l < len(costs) {
				heap.Push(lPQ, costs[l])
			}
			mp[l] = true
			l++
			continue
		}

		if (*lPQ)[0] <= (*rPQ)[0] {
			result += int64(heap.Pop(lPQ).(int))
			if !mp[l] && l < len(costs) {
				heap.Push(lPQ, costs[l])
			}
			mp[l] = true
			l++
		} else {
			result += int64(heap.Pop(rPQ).(int))
			if !mp[r] && r >= 0 {
				heap.Push(rPQ, costs[r])
			}
			mp[r] = true
			r--
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
