package main

import "fmt"

func main() {
	fmt.Println(minJumps([]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}))
}

func minJumps(arr []int) int {
	if len(arr) == 1 {
		return 0
	}

	mp := make(map[int][]int)
	for i, v := range arr {
		mp[v] = append(mp[v], i)
	}

	return bfs(arr, mp)
}

func bfs(arr []int, mp map[int][]int) (result int) {
	q := []int{0}
	visited := make(map[int]bool)
	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			if visited[n] {
				continue
			}
			visited[n] = true

			if n == len(arr)-1 {
				return result
			}

			if n > 0 {
				q = append(q, n-1)
			}
			if n < len(arr) {
				q = append(q, n+1)
			}

			for _, v := range mp[arr[n]] {
				if v != n {
					q = append(q, v)
				}
			}

			mp[arr[n]] = nil
		}
		result++
	}

	return result
}
