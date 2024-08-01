package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(checkArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}))
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	result := make([]bool, len(l))

	for i := 0; i < len(l); i++ {
		tmp := make([]int, r[i]-l[i]+1)
		copy(tmp, nums[l[i]:r[i]+1])

		sort.Ints(tmp)

		isValid := true
		for i := 1; i < len(tmp)-1; i++ {
			if tmp[i+1]-tmp[i] != tmp[1]-tmp[0] {
				isValid = false
				break
			}
		}

		if isValid {
			result[i] = true
		}
	}

	return result
}
