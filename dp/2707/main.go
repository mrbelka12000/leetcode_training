package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minExtraChar("leetscode", []string{"leet", "code", "leetcode"}))
}

func minExtraChar(s string, dictionary []string) int {

	mp := make(map[string]bool)

	for _, v := range dictionary {
		mp[v] = true
	}

	return runner(s, mp, 0, make(map[info]int))
}

type info struct {
	s string
	k int
}

func runner(s string, mp map[string]bool, k int, memo map[info]int) int {
	if len(s) == 0 {
		return k
	}

	key := info{s, k}
	if val, ok := memo[key]; ok {
		return val
	}

	result := math.MaxInt32
	for i := 0; i < len(s); i++ {
		if mp[s[0:i+1]] {
			val := runner(s[i+1:], mp, k, memo)
			result = min(result, val)
		}
	}

	result = min(result, runner(s[1:], mp, k+1, memo))
	memo[key] = result
	return result
}
