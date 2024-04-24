package main

import "fmt"

func main() {
	fmt.Println(validPalindrome("ebcbbececabbacecbbcbe"))
}

func validPalindrome(s string) bool {

	l, r := 0, len(s)-1
	for l < r {
		if s[l] == s[r] {
			l++
			r--
			continue
		}

		if isPalindrome(s[l:r]) {
			return true
		}

		if isPalindrome(s[l+1 : r+1]) {
			return true
		}

		return false
	}

	return true
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
