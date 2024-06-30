package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type MRUQueue struct {
	stack []int
}

func Constructor(n int) MRUQueue {
	stack := make([]int, n)

	for i := 0; i < n; i++ {
		stack[i] = i + 1
	}

	return MRUQueue{
		stack: stack,
	}
}

func (this *MRUQueue) Fetch(k int) int {
	val := this.stack[k-1]

	this.stack = append(this.stack[:k-1], this.stack[k:]...)
	this.stack = append(this.stack, val)

	return val
}

/**
 * Your MRUQueue object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Fetch(k);
 */
