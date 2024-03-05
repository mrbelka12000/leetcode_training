package main

import "fmt"

func main() {
	fmt.Println(numTilePossibilities("ABC"))
}

func numTilePossibilities(tiles string) int {
	if len(tiles) == 1 {
		return 1
	}

	var count int
	visited = make(map[string]struct{})

	runner(tiles, "", &count, map[int]int{})

	return count
}

var visited map[string]struct{}

func runner(tiles, tmp string, count *int, mp map[int]int) {

	if tmp != "" {
		if _, ok := visited[tmp]; ok {
			return
		}
		*count++
	}

	visited[tmp] = struct{}{}

	for i := 0; i < len(tiles); i++ {
		if mp[i] != 0 {
			continue
		}

		mp[i]++
		runner(tiles, tmp+string(tiles[i]), count, mp)
		mp[i]--
	}
}
