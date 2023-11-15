package main

import (
	"fmt"
)

func main() {

	//fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{8, 7, 4, 3}, 11))
}
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int

	solver(0, target, 0, candidates, []int{}, &result)

	return result
}

func solver(ind int, target, curr int, candidates, tmp []int, result *[][]int) bool {
	fmt.Println(tmp, curr)
	if target == curr {
		ans := make([]int, len(tmp))
		copy(ans, tmp)
		// fmt.Println(ans, "append")
		*result = append(*result, ans)
		return true
	}

	if target < curr {
		return false
	}

	for i := ind; i < len(candidates); i++ {
		if !solver(i, target, candidates[i]+curr, candidates, append(tmp, candidates[i]), result) {
			if len(tmp) > 1 && curr > target {
				curr -= tmp[len(tmp)-1]
				tmp = tmp[:len(tmp)-1]
			}

			//fmt.Println(tmp, target)
			//return true
		}
	}

	return true
}
