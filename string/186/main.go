package main

import (
	"strings"
)

func main() {
}

func reverseWords(s []byte) {
	var str strings.Builder

	for i, v := range s {
		str.WriteByte(v)
		s[i] = ' '
	}

	full := str.String()
	arr := strings.Split(full, " ")
	var index int
	for i := len(arr) - 1; i >= 0; i-- {
		str := arr[i]

		for j := 0; j < len(str); j++ {
			s[index] = str[j]
			index++
		}
		index++
	}
}
