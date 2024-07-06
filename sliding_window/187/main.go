package main

import (
	"fmt"
)

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}

func findRepeatedDnaSequences(s string) []string {
	mp := make(map[string]int)
	for i := 0; i <= len(s)-10; i++ {
		mp[s[i:i+10]]++
	}

	var result []string
	for k, v := range mp {
		if v > 1 {
			result = append(result, k)
		}
	}

	return result
}
