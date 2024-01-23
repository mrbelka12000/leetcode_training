package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxLength([]string{"aa", "bb", "ue"}))
	fmt.Println(maxLength([]string{"jnfbyktlrqumowxd", "mvhgcpxnjzrdei"}))

	fmt.Println(maxLength([]string{"cha", "r", "act", "ers"}))
	fmt.Println(maxLength([]string{"abcdefghijklm", "bcdefghijklmn", "cdefghijklmno", "defghijklmnop", "efghijklmnopq", "fghijklmnopqr", "ghijklmnopqrs", "hijklmnopqrst", "ijklmnopqrstu", "jklmnopqrstuv", "klmnopqrstuvw", "lmnopqrstuvwx", "mnopqrstuvwxy", "nopqrstuvwxyz", "opqrstuvwxyza", "pqrstuvwxyzab"}))
}

func maxLength(arr []string) int {
	var result int

	solver(arr, 0, "", &result)

	return result
}

func solver(arr []string, ind int, tmp string, maxVal *int) {
	val, ok := isValid(tmp)
	if ok {
		*maxVal = max(val, *maxVal)
	} else {
		return
	}

	for i := ind; i < len(arr); i++ {
		solver(arr, i+1, tmp+arr[i], maxVal)
	}

}

func isValid(str string) (int, bool) {
	mp := make(map[byte]struct{})

	for i := 0; i < len(str); i++ {
		if _, ok := mp[str[i]]; ok {
			return 0, false
		}
		mp[str[i]] = struct{}{}
	}

	return len(mp), true
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
