package main

import "fmt"

func main() {
	//fmt.Println(countRoutes([]int{1, 2, 3}, 0, 2, 40)) // 4
	fmt.Println(countRoutes([]int{1, 2, 3}, 0, 2, 5)) // 4

}

func countRoutes(locations []int, start int, finish int, fuel int) int {
	if abs(locations[start]-locations[finish]) > fuel {
		return 0
	}
	cache := make([][]int, len(locations))

	for i := 0; i < len(locations); i++ {
		cache[i] = make([]int, fuel+1)
	}

	val := dfs(locations, start, finish, fuel, cache)

	return val
}

const mod = 1000000007

var (
// visited map[[2]int]int
)

func dfs(locations []int, currInd, finishInd, fuel int, cache [][]int) int {
	if fuel < 0 {
		return 0
	}

	if cache[currInd][fuel] > 0 {
		return cache[currInd][fuel]
	}

	if currInd == finishInd {
		cache[currInd][fuel] = 1
	}

	for i, v := range locations {
		if currInd != i && fuel >= abs(locations[currInd]-locations[finishInd]) {
			cache[currInd][fuel] = (cache[currInd][fuel] + dfs(locations, i, finishInd, fuel-abs(locations[currInd]-v), cache)) % mod
		}
	}

	return cache[currInd][fuel]
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
