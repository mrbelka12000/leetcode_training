package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(validWordSquare([]string{"abcd", "bnrt", "crmy", "dtye"}))
}

func validWordSquare(words []string) bool {
	for i := 1; i < len(words); i++ {
		if len(words[i]) > len(words[i-1]) {
			return false
		}
	}

	l, r := 0, 0

	for {
		if l >= len(words) || r >= len(words[l]) {
			break
		}

		var vert, hor strings.Builder
		for i := l; i < len(words) && r < len(words[i]); i++ {
			vert.WriteByte(words[i][r])
		}

		for i := r; i < len(words[l]) && l < len(words); i++ {
			hor.WriteByte(words[l][i])
		}

		// fmt.Println(vert, hor, l , r)

		if vert.String() != hor.String() {
			return false
		}

		l++
		r++
	}

	return true
}

/*

["abcd"
,"befd"
,"cf"
,"dda"]

*/
