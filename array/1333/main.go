package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(filterRestaurants(
		[][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}},
		1,
		50,
		10))
}

func filterRestaurants(restaurants [][]int, vf int, mp int, md int) []int {
	var infos []info

	for _, v := range restaurants {
		if v[2] != vf && vf == 1 {
			continue
		}
		if v[3] > mp {
			continue
		}
		if v[4] > md {
			continue
		}
		infos = append(infos, newInfo(v))
	}

	sort.Slice(infos, func(i, j int) bool {
		if infos[i].rating == infos[j].rating {
			return infos[i].id > infos[j].id
		}
		return infos[i].rating > infos[j].rating
	})
	var result []int

	for _, v := range infos {
		result = append(result, v.id)
	}

	return result
}

func newInfo(arr []int) info {
	return info{
		id:       arr[0],
		rating:   arr[1],
		vegan:    arr[2],
		price:    arr[3],
		distance: arr[4],
	}
}

type info struct {
	id       int
	rating   int
	vegan    int
	price    int
	distance int
}
