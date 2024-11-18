package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfSubstrings("abacb", 2))
}

func numberOfSubstrings(s string, k int) int64 {
	var (
		result int64
		start  int
		arr    [26]int
	)

	for i := 0; i < len(s); i++ {
		char := s[i] - 'a'
		arr[char]++
		for arr[char] == k {
			result += int64(len(s) - i)
			arr[s[start]-'a']--
			start++
		}
	}

	return result
}
