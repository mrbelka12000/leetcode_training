package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello world")
}

func minMeetingRooms(it [][]int) int {
	sort.Slice(it, func(i, j int) bool {
		return it[i][0] < it[j][0]
	})
	rooms := [][]int{it[0]}
	for i := 1; i < len(it); i++ {
		var (
			canMerge bool
		)
		for _, v := range rooms {
			if v[1] <= it[i][0] {
				canMerge = true
				v[1] = it[i][1]
				break
			}
		}
		if !canMerge {
			rooms = append(rooms, it[i])
		}
	}

	return len(rooms)
}
