package main

import (
	"fmt"
)

func main() {
	fmt.Println(validSubstringCount("abcabc", "abc"))
}

func validSubstringCount(word1 string, word2 string) int64 {
	var (
		a      [26]int
		b      [26]int
		result int64
	)
	x := 0
	for _, v := range word2 {
		x++
		b[v-'a']++
	}

	l := 0
	for i, v := range word1 {
		ind := v - 'a'
		a[ind]++

		if a[ind] <= b[ind] {
			x--
		}

		for x == 0 {
			result += int64(len(word1) - i)
			ind := word1[l] - 'a'
			a[ind]--
			l++
			if a[ind] < b[ind] {
				x++
			}
		}
	}

	// fmt.Println(result)
	return result
}
