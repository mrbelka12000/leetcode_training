package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfSpecialSubstrings("abcd"))
}

func numberOfSpecialSubstrings(s string) int {

	var (
		result, start int
		mp            [26]int
	)

	for i := 0; i < len(s); i++ {
		mp[s[i]-'a']++

		for mp[s[i]-'a'] > 1 {
			mp[s[start]-'a']--
			start++
		}

		result += i - start + 1
	}

	return result
}
