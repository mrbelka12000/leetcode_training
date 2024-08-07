package main

import (
	"fmt"
)

func main() {
	fmt.Println(detectCapitalUse("leetcode"))
}

func detectCapitalUse(word string) bool {
	var (
		count int
		check bool
	)
	for i := 0; i < len(word); i++ {
		if word[i] >= 'A' && word[i] <= 'Z' {
			count++
			if check {
				return false
			}
		} else {
			check = true
		}
	}
	if count == 0 {
		return true
	}

	return count > 0 && (count == 1 || count == len(word))
}
