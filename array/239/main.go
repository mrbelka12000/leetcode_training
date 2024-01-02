package main

import "fmt"

func main() {
	fmt.Println(maxSlidingWindow([]int{2, 1, 4, 5, 3, 4, 1, 2}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	var q []int
	var ind []int
	for i := 0; i < len(nums); i++ {

		for len(q) != 0 && q[len(q)-1] < nums[i] {
			q = q[:len(q)-1]
			ind = ind[:len(ind)-1]
		}
		ind = append(ind, i)
		q = append(q, nums[i])

		if i >= k-1 {
			if len(ind) != 0 {
				if i-ind[0] == k {
					ind = ind[1:]
					q = q[1:]
				}
			}
			result = append(result, q[0])
		}

		// fmt.Println(q, ind)
	}

	return result
}
