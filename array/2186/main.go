package main

import "fmt"

func main() {
	fmt.Println(minSteps("leetcode", "coats"))
}

func minSteps(s string, t string) int {
	var arr [26]int
	var arr2 [26]int

	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		arr2[t[i]-'a']++
	}

	var result int

	for i, v := range arr {
		result += abs(v - arr2[i])
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
