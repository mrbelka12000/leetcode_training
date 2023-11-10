package main

import "fmt"

func main() {
	fmt.Println(numSplits("aacaba"))
}

func numSplits(s string) int {
	var result int
	mp1 := map[byte]int{
		s[0]: 1,
	}

	mp2 := make(map[byte]int)

	for i := 1; i < len(s); i++ {
		mp2[s[i]]++
	}

	if len(mp1) == len(mp2) {
		result++
	}

	for i := 1; i < len(s); i++ {
		// fmt.Println(mp1,mp2, s[:i],s[i:])
		val := mp2[s[i]]
		if val == 1 {
			delete(mp2, s[i])
		} else {
			mp2[s[i]] = val - 1
		}
		mp1[s[i]]++
		if len(mp1) == len(mp2) {
			result++
		}
	}

	return result
}
