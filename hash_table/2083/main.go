package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfSubstrings("abcba"))
}

func numberOfSubstrings(s string) int64 {
	var (
		result int
		arr    [26]int
	)

	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
		result += arr[s[i]-'a']
	}

	return int64(result)
}
