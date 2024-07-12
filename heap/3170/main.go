package main

import (
	"container/heap"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(clearStars("aaba*"))
}

func clearStars(s string) string {
	pq := &PQ{}

	heap.Init(pq)

	b := []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			b[i] = '-'
			val := heap.Pop(pq).(check)
			b[val.pos] = '-'
			continue
		}

		heap.Push(pq, check{
			s[i], i,
		})
	}

	result := strings.Builder{}

	for _, v := range b {
		if v == '-' {
			continue
		}
		result.WriteByte(v)
	}

	return result.String()
}

type PQ []check

type check struct {
	b   byte
	pos int
}

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// max-heap
func (pq PQ) Less(i, j int) bool {
	if pq[i].b == pq[j].b {
		return pq[i].pos > pq[j].pos
	}
	return pq[i].b < pq[j].b
}

func (pq *PQ) Push(x interface{}) {
	tmp := x.(check)
	*pq = append(*pq, tmp)
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	tmp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return tmp
}
