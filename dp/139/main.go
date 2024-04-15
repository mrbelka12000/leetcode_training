package main

import "fmt"

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)

	for _, v := range wordDict {
		dict[v] = true
	}
	cache = make(map[string]bool)
	return runner(s, dict)
}

var cache map[string]bool

func runner(s string, dict map[string]bool) bool {
	if s == "" {
		return true
	}

	if cache[s] {
		return false
	}
	cache[s] = true

	for i := 0; i < len(s); i++ {
		if dict[s[:i+1]] {
			if runner(s[i+1:], dict) {
				return true
			}
		}
	}

	return false
}
