package main

import "fmt"

func main() {
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5))
}

func canReach(arr []int, start int) bool {
	return bfs(arr, start)
}

func bfs(arr []int, start int) bool {
	q := []int{start}
	mp := make(map[int]struct{})

	for len(q) != 0 {
		pos := q[0]
		q = q[1:]
		if arr[pos] == 0 {
			return true
		}
		mp[pos] = struct{}{}

		l, r := pos-arr[pos], pos+arr[pos]
		if l >= 0 {
			if _, ok := mp[l]; !ok {
				q = append(q, l)
			}
		}
		if r < len(arr) {
			if _, ok := mp[r]; !ok {
				q = append(q, r)
			}
		}
	}

	return false
}
