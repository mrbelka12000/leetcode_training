package main

import "fmt"

func main() {
	fmt.Println(numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6))

}

func numBusesToDestination(routes [][]int, source int, target int) int {
	busCons := make(map[int]map[int]bool)
	routeCons := make(map[int][]int)

	for i := 0; i < len(routes); i++ {
		busCons[i] = make(map[int]bool)
		for j := 0; j < len(routes[i]); j++ {

			busCons[i][routes[i][j]] = true

			routeCons[routes[i][j]] = append(routeCons[routes[i][j]], i)
		}
	}

	fmt.Println(busCons)
	fmt.Println(routeCons)
	return runner(busCons, routeCons, source, target)
}

func runner(busCons map[int]map[int]bool, routeCons map[int][]int, start, end int) int {
	q := []int{start}

	var dist int
	visited := make(map[int]bool)
	for len(q) > 0 {

		size := len(q)

		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]
			if n == end {
				return dist
			}

			for _, v := range routeCons[n] {
				if visited[v] {
					continue
				}
				for k := range busCons[v] {
					if k != n {
						q = append(q, k)
					}
				}
				visited[v] = true
			}
		}

		dist++
	}

	fmt.Println(visited)
	return -1
}

/*
[[1,2,7]
,[3,6,7]]
*/
