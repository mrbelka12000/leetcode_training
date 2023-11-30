package main

import "math/rand"

func main() {

}

type Solution struct {
	mp map[int][]int
}

func Constructor(nums []int) Solution {
	s := Solution{
		mp: make(map[int][]int),
	}

	for i, v := range nums {
		s.mp[v] = append(s.mp[v], i)
	}

	return s
}

func (this *Solution) Pick(target int) int {
	arr := this.mp[target]

	return arr[rand.Intn(len(arr))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
