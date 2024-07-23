package main

import (
	"fmt"
)

func main() {
	fmt.Println(partition("aab"))
}

func partition(s string) [][]string {
	result = nil

	runner(s, nil)

	return result
}

var result [][]string

func runner(s string, curr []string) {
	if s == "" {
		tmp := make([]string, len(curr))
		copy(tmp, curr)
		result = append(result, tmp)
		return
	}

	for i := 0; i < len(s); i++ {
		if isPalindrome(s[:i+1]) {
			runner(s[i+1:], append(curr, s[:i+1]))
		}
	}
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
