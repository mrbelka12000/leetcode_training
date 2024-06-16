package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type NumArray struct {
	bit []int
}

func Constructor(nums []int) NumArray {
	bit := make([]int, len(nums)+1)
	for i := 1; i < len(bit); i++ {
		bit[i] += nums[i-1]
		j := lsb(i)
		if j+i < len(bit) {
			bit[j+i] += bit[i]
		}
	}

	return NumArray{bit}
}

func (this *NumArray) Update(index int, val int) {
	d := val - this.SumRange(index, index)
	for i := index + 1; i < len(this.bit); i += lsb(i) {
		this.bit[i] += d
	}
}

func (this *NumArray) sum(i int) int {
	var sum int
	for i > 0 {
		sum += this.bit[i]
		i -= lsb(i)
	}
	return sum
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum(right+1) - this.sum(left)
}

func lsb(i int) int {
	return i & -i
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
