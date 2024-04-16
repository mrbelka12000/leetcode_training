package main

import "fmt"

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))

}

func findSubstring(s string, words []string) []int {
	dic := make(map[string]int)
	for _, v := range words {
		dic[v]++
	}

	total = len(words)

	var (
		result []int
		k      = len(words[0])
	)

	for i := 0; i < len(s)-len(words)*k+1; i++ {
		key := s[i : i+k]
		if _, ok := dic[key]; ok {
			dic[key]--
			if runner(s[i+k:], dic, k, 1) {
				result = append(result, i)
			}
			dic[key]++
		}
	}

	return result
}

var total int

func runner(s string, dic map[string]int, k int, count int) bool {

	if count == total {
		return true
	}

	if len(s) < k*(total-count) {
		return false
	}

	key := s[:k]

	if val := dic[key]; val != 0 {
		dic[key]--
		check := runner(s[k:], dic, k, count+1)
		dic[key]++
		return check
	}

	return false
}
