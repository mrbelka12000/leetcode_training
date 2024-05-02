package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minDeletions("aaabbbcc"))
}

func minDeletions(s string) int {
	var freq [26]int

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	sort.Ints(freq[:])

	var result int

	// fmt.Println(freq)

	for i := len(freq) - 1; i >= 1; i-- {
		if freq[i] == 0 {
			break
		}

		for j := i - 1; j >= 0; j-- {
			if freq[j] == freq[i] {
				freq[j]--
				result++
			}
		}
		// fmt.Println(freq, "in loop", i)
	}

	return result
}
