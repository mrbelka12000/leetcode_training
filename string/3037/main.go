package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type InfiniteStream interface {
	Next() int
}

func findPattern(stream InfiniteStream, pt []int) int {
	pattern := make([]byte, len(pt))

	for i := 0; i < len(pt); i++ {
		pattern[i] = byte(pt[i]) + '0'
	}

	var (
		result int
		ind    int
		q      []byte
	)

	for ind != len(pattern) {
		val := getByte(stream.Next())
		if val == pattern[ind] {
			q = append(q, val)
			ind++
		} else {
			if ind > 0 {
				q = append(q, val)
				// fmt.Println(string(pattern[:ind+1]), string(q), "before")
				for len(q) > 0 && string(q) != string(pattern[:ind+1]) {
					ind--
					q = q[1:]
				}
				ind++
				// fmt.Println(string(pattern[:ind+1]), string(q), "after")
			}
		}
		result++
		// fmt.Printf("ind = %v , result = %v, val = %v, q = %v\n", ind, result, string(val), string(q))
	}
	// fmt.Println(result, ind)
	return result - ind
}

func getByte(i int) byte {
	return byte(i) + '0'
}
