package main

func main() {

}

func canPartition(nums []int) bool {
	var sum int

	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	memo = make(map[[2]int]bool)

	return runner(nums, len(nums)-1, target)
}

var memo map[[2]int]bool

func runner(nums []int, n, target int) bool {
	if target == 0 {
		return true
	}

	if n == 0 || target < 0 {
		return false
	}

	if val, ok := memo[[2]int{n, target}]; ok {
		return val
	}

	if nums[n-1] > target {
		return runner(nums, n-1, target)
	}

	memo[[2]int{n, target}] = runner(nums, n-1, target-nums[n-1]) || runner(nums, n-1, target)

	return memo[[2]int{n, target}]
}
