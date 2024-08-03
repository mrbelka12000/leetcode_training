package main

import (
	"fmt"
)

func main() {
	fmt.Println(characterReplacement("AABABBA", 1))
}

func characterReplacement(str string, k int) int {
	var (
		result, start int
		mp            = make(map[byte]int)
	)

	for i := 0; i < len(str); i++ {
		mp[str[i]]++
		for !isValid(mp, k) {
			mp[str[start]]--
			if mp[str[start]] == 0 {
				delete(mp, str[start])
			}
			start++
		}
		result = max(result, i-start+1)
	}

	return result
}

func isValid(mp map[byte]int, k int) bool {
	var (
		maxCount int
	)

	for _, v := range mp {
		maxCount = max(maxCount, v)
	}

	k += maxCount

	for _, v := range mp {
		k -= v

		if k < 0 {
			return false
		}

	}

	return true
}
