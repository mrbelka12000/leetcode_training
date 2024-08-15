package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMaximizedCapital(1, 2, []int{1, 2, 3}, []int{1, 1, 2}))
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	store := make([][]int, len(profits))

	for i := 0; i < len(profits); i++ {
		store[i] = []int{capital[i], profits[i]}
	}
	sort.Slice(store, func(i, j int) bool {
		return store[i][0] < store[j][0]
	})

	start := getStart(store, w)
	if start == -1 {
		return w
	}

	pq := PQ(store[:start+1])
	heap.Init(&pq)
	store = store[start+1:]

	result := w

	for k > 0 && pq.Len() > 0 {
		item := heap.Pop(&pq).([]int)
		result += item[1]

		if len(store) > 0 {
			start := getStart(store, result)
			if start != -1 {
				// fmt.Println("popal", k, store[:10], result)
				for _, v := range store[:start+1] {
					heap.Push(&pq, v)
				}
				store = store[start+1:]
			}
		}
		k--
	}

	return result
}

func getStart(store [][]int, w int) int {
	l, r := -1, len(store)

	for r-l > 1 {
		mid := (l + r) / 2

		if store[mid][0] > w {
			r = mid
		} else {
			l = mid
		}
	}

	return l
}

type PQ [][]int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// max-heap
func (pq PQ) Less(i, j int) bool {
	return pq[i][1] > pq[j][1]
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
