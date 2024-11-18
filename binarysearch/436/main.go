package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findRightInterval([][]int{{3, 4}, {2, 3}, {1, 2}}))
}

func findRightInterval(it [][]int) []int {
	var infos []info
	for i, v := range it {
		infos = append(infos, info{
			arr: v,
			ind: i,
		})
	}

	sort.Slice(infos, func(i, j int) bool {
		return infos[i].arr[0] < infos[j].arr[0]
	})

	result := make([]int, len(it))
	// fmt.Println(infos)
	for i, v := range it {
		result[i] = findRight(infos, v)
	}

	return result
}

// [1,2],[2,3],[3,4]
func findRight(infos []info, target []int) int {
	// fmt.Println(target)
	l, r := -1, len(infos)
	var found bool
	for r-l > 1 {

		mid := l + (r-l)/2

		if infos[mid].arr[0] >= target[1] {
			r = mid
			found = true
		} else {
			l = mid
		}
	}
	if !found {
		return -1
	}

	return infos[r].ind
}

type info struct {
	arr []int
	ind int
}
