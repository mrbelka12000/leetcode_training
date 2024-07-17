package main

import (
	"fmt"
)

func main() {
	fmt.Println(taskSchedulerII([]int{1, 2, 1, 2, 3, 1}, 3))
}

func taskSchedulerII(tasks []int, space int) int64 {
	var result int
	mp := make(map[int]int)
	for _, v := range tasks {
		result++
		// fmt.Println(mp, v)
		if val, ok := mp[v]; ok {
			// fmt.Println(val, space, result)
			if val+space >= result {
				result = val + space + 1
			}
			mp[v] = result
			continue
		}

		mp[v] = result
	}

	// fmt.Println(mp)
	return int64(result)
}
