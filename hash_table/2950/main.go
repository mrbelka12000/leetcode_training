package main

import (
	"fmt"
)

func main() {
	fmt.Println(countDivisibleSubstrings("abcd"))
}

func countDivisibleSubstrings(word string) int {
	var result int
	var arr = [26]int{1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 05, 6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9}

	for i := 0; i < len(word); i++ {
		var sum int
		for j := i; j < len(word); j++ {
			sum += arr[word[j]-'a']
			if sum%(j-i+1) == 0 {
				result++
			}
		}
	}

	return result
}
