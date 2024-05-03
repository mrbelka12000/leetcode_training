package main

func main() {

}

func returnToBoundaryCount(nums []int) int {
	pref := make([]int, len(nums))
	pref[0] = nums[0]

	var result int

	for i := 1; i < len(nums); i++ {
		pref[i] = pref[i-1] + nums[i]

		if pref[i] == 0 {
			result++
			pref[i] = 0
		}
	}

	return result
}
