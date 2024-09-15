package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countGreatEnoughNodes(root *TreeNode, k int) int {

	var (
		result int
		runner func(root *TreeNode) *PQ
	)

	runner = func(node *TreeNode) *PQ {
		if node == nil {
			return nil
		}

		pq := &PQ{}
		heap.Init(pq)
		l := runner(node.Left)
		r := runner(node.Right)
		if l != nil {
			for l.Len() > 0 {
				heap.Push(pq, heap.Pop(l).(int))
				if pq.Len() > k {
					heap.Pop(pq)
				}
			}
		}

		if r != nil {
			for r.Len() > 0 {
				heap.Push(pq, heap.Pop(r).(int))
				if pq.Len() > k {
					heap.Pop(pq)
				}
			}
		}

		if pq.Len() != 0 {
			val := heap.Pop(pq).(int)

			if val < node.Val && pq.Len()+1 >= k {
				result++
			}
			heap.Push(pq, val)
		}
		heap.Push(pq, node.Val)

		return pq
	}

	runner(root)
	return result
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
