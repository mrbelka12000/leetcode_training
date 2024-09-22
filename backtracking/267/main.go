package main

import (
	"fmt"
)

func main() {
	fmt.Println(generatePalindromes("aabbhijkkjih"))
}

func generatePalindromes(s string) []string {
	result = nil
	var tmp []byte
	for i := 0; i < len(s); i++ {
		tmp = append(tmp, '1')
	}
	runner(s, tmp, make([]bool, len(s)), 0, len(s)-1, true)

	return result
}

var (
	result []string
)

func runner(s string, tmp []byte, used []bool, k, last int, dir bool) {
	if isPalindrome(string(tmp)) {
		result = append(result, string(tmp))
	}

	check := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if used[i] || check[s[i]] {
			continue
		}
		check[s[i]] = true
		used[i] = true

		prev := tmp[k]
		if dir {
			tmp[k] = s[i]
			runner(s, tmp, used, k+1, last, !dir)
		} else {
			if tmp[k-1] == s[i] {
				tmp[last] = s[i]
				runner(s, tmp, used, k, last-1, !dir)
			}
		}

		tmp[k] = prev
		used[i] = false
	}
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] == '1' {
			return false
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
