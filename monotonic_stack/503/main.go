package main

func main() {

}

func nextGreaterElements(nums []int) []int {
	ans := make([]int, len(nums))
	var max int = -1234123412341

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		var (
			mid   = len(nums) / 2
			check bool
		)

		if nums[i] == max {
			ans[i] = -1
			continue
		}

		for j := i + 1; j <= len(nums); j++ {
			if j == len(nums) {
				j = 0
			}

			if j == mid {
				if check {
					break
				}
				check = true
			}

			if nums[j] > nums[i] {
				ans[i] = nums[j]
				break
			}
		}
	}

	return ans
}
