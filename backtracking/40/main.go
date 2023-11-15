package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(combinationSum2([]int{3, 1, 3, 5, 1, 1}, 8))
	//fmt.Println(combinationSum2([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 30))
}

func combinationSum2(candidates []int, target int) [][]int {
	check := make(map[int]struct{})

	for _, v := range candidates {
		check[v] = struct{}{}
	}

	mp = make(map[string]struct{})
	sort.Ints(candidates)
	var result [][]int

	solver(len(check), 0, target, 0, candidates, []int{}, &result)

	return result
}

var mp map[string]struct{}

func solver(max, ind, target, curr int, candidates, tmp []int, result *[][]int) bool {
	//fmt.Println(mp)

	if target == curr {

		strKey := fmt.Sprint(tmp)
		_, ok := mp[strKey]
		if !ok {
			ans := make([]int, len(tmp))
			copy(ans, tmp)
			*result = append(*result, ans)

			mp[strKey] = struct{}{}
		}

		return true
	}

	if target < curr {
		return false
	}

	for i := ind; i < len(candidates); i++ {
		if (max == 1 || max == 2) && len(*result) == max {
			return false
		}
		if !solver(max, i+1, target, candidates[i]+curr, candidates, append(tmp, candidates[i]), result) {
			if len(tmp) > 1 && curr > target {
				curr -= tmp[len(tmp)-1]
				tmp = tmp[:len(tmp)-1]
			}
		}
	}

	return true
}
