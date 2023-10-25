package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	var result string

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			str := string(s[i : j+1])
			// fmt.Println(str)
			if isPalindrome(str) {
				if len(result) < len(str) {
					result = str
				}
			}
		}
	}

	return result
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
