package main

import "math/rand"

func main() {

}

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	var arr []int

	mp := make(map[int]struct{})
	for len(arr) != len(this.nums) {
		n := rand.Intn(len(this.nums))
		if _, ok := mp[n]; ok {
			continue
		}
		mp[n] = struct{}{}
		arr = append(arr, this.nums[n])
	}

	return arr
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
