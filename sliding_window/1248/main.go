package main

func main() {

}

func numberOfSubarrays(nums []int, k int) int {
	return find(nums, k) - find(nums, k-1)
}

func find(nums []int, k int) int {
	var start, result int

	for ind, v := range nums {
		k -= v % 2

		for k < 0 {
			k += nums[start] % 2
			start++
		}

		result += ind - start + 1
	}

	return result
}
