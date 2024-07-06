package main

import (
	"fmt"
)

func main() {
	fmt.Println(numKLenSubstrNoRepeats("havefunonleetcode", 5))
}

func numKLenSubstrNoRepeats(s string, k int) int {
	var start, result int

	mp := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		mp[s[i]]++

		for len(mp) > k || mp[s[i]] > 1 {
			mp[s[start]]--
			if mp[s[start]] == 0 {
				delete(mp, s[start])
			}
			start++
		}

		if len(mp) == k {
			result++
		}
	}

	return result
}
