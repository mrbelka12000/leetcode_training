package main

import "sort"

func main() {

}

func findMinArrowShots(p [][]int) int {
	if len(p) == 1 {
		return 1
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i][0] < p[j][0]
	})

	sort.Slice(p, func(i, j int) bool {
		return p[i][1] < p[j][1]
	})

	// fmt.Println(p)

	var (
		result int
		check  int
	)

	for i := 0; i < len(p); i++ {
		var blow bool
		for j := i + 1; j < len(p); j++ {
			if p[i][1] >= p[j][0] {
				// fmt.Println(p[i],p[j],"inseide")
				blow = true
				check = j
				continue
			}
			break
		}
		if blow {
			i = check
		}
		result++
	}

	return result
}
