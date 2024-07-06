package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func lengthOfLongestSubstringTwoDistinct(s string) int {
	mp := make(map[byte]int)

	var start, result int

	for i := 0; i < len(s); i++ {
		mp[s[i]]++
		for len(mp) > 2 {
			mp[s[start]]--
			if mp[s[start]] == 0 {
				delete(mp, s[start])
			}
			start++
		}
		result = max(result, i-start+1)
	}

	return result
}
