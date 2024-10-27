package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(containVirus([][]int{
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}))
}

func containVirus(isInfected [][]int) int {
	var result int

	infos := getInfos(isInfected)

	for len(infos) > 0 {
		sort.Slice(infos, func(i, j int) bool {
			return len(infos[i].mp) > len(infos[j].mp)
		})

		// for _, v := range infos {
		// 	fmt.Printf("%+v\n", v)
		// }
		// fmt.Println()

		result += infos[0].Perimeter

		for _, area := range infos[0].Area {
			isInfected[area[0]][area[1]] = -1
		}

		visited := make([][]bool, len(isInfected))
		for i := range visited {
			visited[i] = make([]bool, len(isInfected[i]))
		}

		for _, area := range infos[1:] {
			for _, v := range area.Area {
				extendZone(isInfected, visited, v[0], v[1])
			}
		}

		infos = getInfos(isInfected)

	}
	return result
}

type Info struct {
	Area      [][]int
	Perimeter int
	Count     int
	mp        map[[2]int]bool
}

func getInfos(isInfected [][]int) []Info {
	var (
		infos   []Info
		visited = make([][]bool, len(isInfected))
	)
	for i := 0; i < len(isInfected); i++ {
		visited[i] = make([]bool, len(isInfected[i]))
	}

	for i := 0; i < len(isInfected); i++ {
		for j := 0; j < len(isInfected[i]); j++ {
			if isInfected[i][j] == 1 && !visited[i][j] {
				in := &Info{
					Area: [][]int{{i, j}},
					mp:   make(map[[2]int]bool),
				}
				visited[i][j] = true
				pem := in.getPerimeter(isInfected, visited, i, j)
				in.Perimeter = pem
				infos = append(infos, *in)
			}
		}
	}

	return infos
}

func extendZone(zone [][]int, visited [][]bool, x, y int) {
	if x < 0 || x >= len(zone) || y < 0 || y >= len(zone[0]) || visited[x][y] {
		return
	}
	visited[x][y] = true
	if zone[x][y] == 0 {
		zone[x][y] = 1
		return
	}

	if zone[x][y] == -1 {
		return
	}

	extendZone(zone, visited, x, y-1)
	extendZone(zone, visited, x, y+1)
	extendZone(zone, visited, x+1, y)
	extendZone(zone, visited, x-1, y)
}

var dirs = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func (i *Info) getPerimeter(zone [][]int, visited [][]bool, x, y int) int {

	var pem int

	for _, dir := range dirs {
		xx, yy := x+dir[0], y+dir[1]
		if xx >= 0 && xx < len(zone) && yy >= 0 && yy < len(zone[0]) && !visited[xx][yy] {
			if zone[xx][yy] == 0 {
				i.mp[[2]int{xx, yy}] = true
				pem++
			} else if zone[xx][yy] == 1 {
				visited[xx][yy] = true
				pem += i.getPerimeter(zone, visited, xx, yy)
				i.Area = append(i.Area, []int{xx, yy})
			}
		}
	}

	return pem
}
