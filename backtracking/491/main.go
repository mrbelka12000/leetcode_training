package main

import "fmt"

func main() {

}

func findSubsequences(nums []int) [][]int {
	mp = make(map[string]struct{})
	var result [][]int

	helper(nums, nil, 0, &result)

	return result
}

var mp = map[string]struct{}{}

func helper(nums, tmp []int, ind int, result *[][]int) {
	if len(tmp) >= 2 {
		if tmp[len(tmp)-1] >= tmp[len(tmp)-2] {
			if _, ok := mp[fmt.Sprint(tmp)]; !ok {
				cp := make([]int, len(tmp))
				copy(cp, tmp)
				*result = append(*result, cp)
				mp[fmt.Sprint(cp)] = struct{}{}
			}
		} else {
			return
		}
	}

	for i := ind; i < len(nums); i++ {
		helper(nums, append(tmp, nums[i]), i+1, result)
	}
}
