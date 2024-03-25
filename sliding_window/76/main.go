package main

import "fmt"

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	arr := [58]int{}

	for i := 0; i < len(t); i++ {
		arr[t[i]-'A']++
	}

	tmp := [58]int{}
	var (
		start  int
		result string
	)
	for i := 0; i < len(s); i++ {
		tmp[s[i]-'A']++

		for isSubstring(arr, tmp) {
			if result == "" || len(result) > i-start+1 {
				result = string(s[start : i+1])
			}
			// fmt.Println(start,result)
			tmp[s[start]-'A']--
			start++
		}
	}

	return result
}

func isSubstring(arr, tmp [58]int) bool {

	for i := 0; i < len(tmp); i++ {
		if arr[i] > tmp[i] {
			return false
		}
	}

	return true
}
