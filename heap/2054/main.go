package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxTwoEvents([][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}))
	fmt.Println(maxTwoEvents([][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}}))
}

func maxTwoEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] == events[j][0] {
			return events[i][1] < events[j][1]
		}
		return events[i][0] < events[j][0]
	})

	suffArray := make([]int, len(events))
	n := len(events)
	suffArray[n-1] = events[n-1][2]
	for i := n - 2; i >= 0; i-- {
		suffArray[i] = max(suffArray[i+1], events[i][2])
	}

	var result int
	// fmt.Println(bSearch(events, 3, 38))
	// fmt.Println(events)
	// fmt.Println(suffArray)

	for i := 0; i < n; i++ {
		optimalPos := bSearch(events, i, events[i][1])
		// fmt.Println(i, optimalPos, events[i])
		if optimalPos < n && events[optimalPos][0] > events[i][1] {
			result = max(result, suffArray[optimalPos]+events[i][2])
			// fmt.Println("sum: ", suffArray[optimalPos]+events[i][2])
		}
		result = max(result, events[i][2])
	}

	return result
}

func bSearch(events [][]int, ind, currE int) int {

	l, r := ind, len(events)

	for r-l > 1 {
		mid := l + (r-l)/2
		if events[mid][0] <= currE {
			l = mid
		} else {
			r = mid
		}
	}

	return r
}
