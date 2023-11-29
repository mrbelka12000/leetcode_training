package main

import "fmt"

func main() {
	fmt.Println(queensAttacktheKing([][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}}, []int{0, 0}))
}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	check := make(map[[2]int]struct{})
	for _, v := range queens {
		check[[2]int{v[0], v[1]}] = struct{}{}
	}

	var result [][]int

	for i := 0; i < len(dirs); i++ {
		find(check, king[0], king[1], i, &result)
	}

	return result
}

var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func find(check map[[2]int]struct{}, x, y, dir int, result *[][]int) {
	if x < 0 || x > 8 || y < 0 || y > 8 {
		return
	}
	if _, ok := check[[2]int{x, y}]; ok {
		*result = append(*result, []int{x, y})
		return
	}

	find(check, x+dirs[dir][0], y+dirs[dir][1], dir, result)
}
