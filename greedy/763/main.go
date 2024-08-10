package main

import (
	"fmt"
)

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}

func partitionLabels(s string) []int {
	mp := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		mp[s[i]] = i
	}

	var result []int

	var maxPart, last int
	for i := 0; i < len(s); i++ {
		maxPart = max(maxPart, mp[s[i]])

		if i == maxPart {
			result = append(result, maxPart-last+1)
			last = maxPart + 1
			maxPart = 0
		}
	}

	return result
}
