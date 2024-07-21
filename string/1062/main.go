package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestRepeatingSubstring("aabcaabdaab"))
}

func longestRepeatingSubstring(s string) int {
	mp := make(map[string]int)

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			mp[s[i:j+1]]++
		}
	}
	var result int

	for k, v := range mp {
		if v > 1 {
			result = max(result, len(k))
		}
	}

	return result
}
