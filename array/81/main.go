package main

import "fmt"

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
}

func search(nums []int, target int) bool {
	if len(nums) == 1 {
		return nums[0] == target
	}

	var ind int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			nums = append(nums, nums[:i+1]...)
			ind = i + 1
			break
		}
	}
	nums = nums[ind:]

	return bSearch(nums, target)
}

func bSearch(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	mp := make(map[int]struct{})
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			r = mid - 1
		} else {
			l = mid + 1
		}
		mp[nums[mid]] = struct{}{}
	}

	return false
}
