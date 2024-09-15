package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello world")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestKValues(root *TreeNode, target float64, k int) []int {
	pq := &PQ{}
	heap.Init(pq)

	var runner func(node *TreeNode)
	runner = func(node *TreeNode) {
		if node == nil {
			return
		}

		heap.Push(pq, info{getDist(target, node.Val), node.Val})
		if pq.Len() > k {
			heap.Pop(pq)
		}
		runner(node.Left)
		runner(node.Right)
	}

	runner(root)

	var result []int

	for pq.Len() > 0 {
		i := pq.Pop().(info)
		result = append(result, i.val)
	}
	return result
}

func getDist(orig float64, x int) float64 {
	return abs(orig - float64(x))
}

func abs(a float64) float64 {
	return math.Abs(a)
}

type info struct {
	dist float64
	val  int
}
type PQ []info

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// max-heap
func (pq PQ) Less(i, j int) bool {
	return pq[i].dist > pq[j].dist
}

func (pq *PQ) Push(x interface{}) {
	tmp := x.(info)
	*pq = append(*pq, tmp)
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	tmp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return tmp
}
