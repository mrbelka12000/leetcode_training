package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello world")
}

func canAttendMeetings(it [][]int) bool {
	sort.Slice(it, func(i, j int) bool {
		return it[i][0] < it[j][1]
	})

	for i := 0; i < len(it)-1; i++ {
		if it[i][1] > it[i+1][0] {
			return false
		}
	}

	return true
}
