package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type MovingAverage struct {
	val  int
	q    []int
	size int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
		q:    make([]int, 0, size+1),
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.q = append(this.q, val)
	this.val += val

	if len(this.q) > this.size {
		sub := this.q[0]
		this.q = this.q[1:]
		this.val -= sub
	}

	return float64(this.val) / float64(len(this.q))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
