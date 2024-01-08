package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 3))
}
func topKFrequent(words []string, k int) []string {
	Map := make(map[string]int)
	for _, word := range words {
		Map[word]++
	}

	pq := &PQ{}
	heap.Init(pq)
	for w, c := range Map {
		heap.Push(pq, &WordCnt{w, c})
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}

	ans := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		wc := heap.Pop(pq).(*WordCnt)
		ans[i] = wc.word
	}
	return ans
}

type WordCnt struct {
	word string
	cnt  int
}

type PQ []*WordCnt

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// min-heap
func (pq PQ) Less(i, j int) bool {
	if pq[i].cnt == pq[j].cnt {
		return pq[i].word > pq[j].word
	}
	return pq[i].cnt < pq[j].cnt
}

func (pq *PQ) Push(x interface{}) {
	tmp := x.(*WordCnt)
	*pq = append(*pq, tmp)
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	tmp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return tmp
}
