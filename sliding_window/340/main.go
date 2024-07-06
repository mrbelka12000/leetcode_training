package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstringKDistinct("eceba", 2))
}

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	mp := make(map[byte]int)

	var start, result int

	for i := 0; i < len(s); i++ {
		mp[s[i]]++
		for len(mp) > k {
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
