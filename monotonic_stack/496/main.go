package main

func main() {

}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	pos := make(map[int]int)
	var stack []int

	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && stack[len(stack)-1] < nums2[i] {
			pos[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, nums2[i])
	}

	for i := 0; i < len(nums1); i++ {
		if val, ok := pos[nums1[i]]; ok {
			ans[i] = val
		} else {
			ans[i] = -1
		}
	}

	return ans
}
