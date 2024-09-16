package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(findMinDifference([]string{"00:00", "23:59", "00:00"}))
}

func findMinDifference(timePoints []string) int {
	var minutes []int

	for _, v := range timePoints {
		minute := int(v[len(v)-1] - '0')
		minute = minute + int(v[len(v)-2]-'0')*10

		h := int(v[len(v)-4] - '0')
		h = h + int(v[len(v)-5]-'0')*10
		minutes = append(minutes, minute+(h*60))
	}
	sort.Ints(minutes)

	result := math.MaxInt32
	l, r := 0, len(minutes)-1
	for l < r {
		curr := getDist(minutes[l], minutes[r])
		result = min(result, curr)

		// fmt.Println(getDist(minutes[l], minutes[l+1]),getDist(minutes[r-1], minutes[r]))
		if getDist(minutes[l], minutes[l+1]) < getDist(minutes[r-1], minutes[r]) {
			r--
		} else {
			l++
		}
	}

	return result
}

func getDist(a, b int) int {
	curr := min(1440, b-a)
	curr = min(curr, 1440-b+a)
	return curr
}
