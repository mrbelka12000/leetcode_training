package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isPossibleDivide([]int{1, 2, 3, 3, 4, 4, 5, 6}, 4))

}

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}
	sort.Ints(nums)

	for len(nums) != 0 {
		var (
			prev  int
			count int
		)
		for i := 0; i < len(nums); i++ {
			if nums[i] == prev {
				continue
			}
			if count == k {
				count = 0
				break
			}
			if prev != 0 {
				if prev+1 != nums[i] && prev-1 != nums[i] {
					continue
				}
			}
			count++
			prev = nums[i]
			nums = remove(nums, i)
			i--
		}
		if count != k && count != 0 {
			break
		}
	}

	return len(nums) == 0
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
