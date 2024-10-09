package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumLengthSubstring("bcbbbcba"))
}

func maximumLengthSubstring(s string) int {
	var (
		arr    [26]int
		result int
		start  int
	)

	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++

		for arr[s[i]-'a'] > 2 {
			arr[s[start]-'a']--
			start++
		}

		result = max(result, i-start+1)
	}

	return result
}
