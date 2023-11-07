package main

import "fmt"

func main() {
	fmt.Println(numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"}))
}

func numMatchingSubseq(s string, words []string) int {
	var result int
	for _, word := range words {
		var ind int
		for i := 0; i < len(s); i++ {
			// fmt.Println(i, ind, len(s), len(word))
			if s[i] == word[ind] {
				ind++
			}

			if ind == len(word) {
				result++
				break
			}
		}
	}

	return result
}
