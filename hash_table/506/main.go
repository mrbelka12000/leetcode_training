package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))
}

func findRelativeRanks(score []int) []string {
	cp := make([]int, len(score))
	copy(cp, score)
	sort.Slice(cp, func(i, j int) bool {
		return cp[i] > cp[j]
	})

	mp := make(map[int]string)
	for _, v := range cp {
		if len(mp) == 0 {
			mp[v] = "Gold Medal"
		} else if len(mp) == 1 {
			mp[v] = "Silver Medal"
		} else if len(mp) == 2 {
			mp[v] = "Bronze Medal"
		} else {
			mp[v] = fmt.Sprint(len(mp) + 1)
		}
	}

	result := make([]string, len(score))
	for i, v := range score {
		if val, ok := mp[v]; ok {
			result[i] = val
		}
	}

	return result
}
