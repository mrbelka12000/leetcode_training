package main

import (
	"fmt"
)

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}

func restoreIpAddresses(s string) []string {
	result = nil

	runner(s, "", 4)

	return result
}

var result []string

func runner(s, tmp string, k int) {
	if k == 0 {
		if len(s) == 0 {
			result = append(result, tmp[:len(tmp)-1])
		}
		return
	}

	var num int
	for i := 0; i < 3; i++ {
		if i >= len(s) {
			return
		}

		if i != 0 {
			if s[0] == '0' {
				return
			}
		}

		num = num*10 + int(s[i]-'0')
		if num > 255 {
			return
		}
		runner(s[i+1:], tmp+s[:i+1]+".", k-1)
	}
}
