package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkValidString("**************************************************))))))))))))))))))))))))))))))))))))))))))))))))))"))
}

func checkValidString(s string) bool {
	memo = make(map[info]bool)
	return runner(s, "", 0)
}

type info struct {
	tmp string
	pos int
}

var memo map[info]bool

func runner(s, tmp string, pos int) bool {
	if pos == len(s) {
		if tmp == "" {
			return true
		}
		return false
	}
	if memo[info{tmp, pos}] {
		return false
	}
	memo[info{tmp, pos}] = true

	if s[pos] == '(' {
		return runner(s, tmp+"(", pos+1)
	} else if s[pos] == ')' {
		if len(tmp) == 0 || tmp[len(tmp)-1] != '(' {
			return false
		}
		return runner(s, tmp[:len(tmp)-1], pos+1)
	} else {
		if len(tmp) != 0 && tmp[len(tmp)-1] == '(' {
			if runner(s, tmp[:len(tmp)-1], pos+1) {
				return true
			}
		}

		return runner(s, tmp, pos+1) || runner(s, tmp+"(", pos+1)
	}
}
