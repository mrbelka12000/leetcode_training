package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type SparseVector struct {
	nums []int
}

func Constructor(nums []int) SparseVector {
	return SparseVector{
		nums: nums,
	}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	var result int

	for i, v := range vec.nums {
		result += (v * this.nums[i])
	}

	return result
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */
