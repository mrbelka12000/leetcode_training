package main

import "fmt"

func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}

func wordBreak(s string, wordDict []string) []string {
	dict := make(map[string]bool)

	for _, v := range wordDict {
		dict[v] = true
	}
	var result []string

	bfs(s, "", dict, &result)

	return result
}

func bfs(s, tmp string, dict map[string]bool, result *[]string) {
	if len(s) == 0 {
		*result = append(*result, tmp[:len(tmp)-1])
		return
	}

	for i := 0; i < len(s); i++ {
		if dict[s[:i+1]] {
			bfs(s[i+1:], tmp+s[:i+1]+" ", dict, result)
		}
	}
}
