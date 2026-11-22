package main

import (
	"fmt"
)

func main() {
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Shogun", "Tapioca Express"}))
}

func findRestaurant(list1 []string, list2 []string) []string {
	mp := make(map[string][3]int)

	for i, v := range list1 {
		mp[v] = [3]int{i, 0, 1}
	}

	for i, v := range list2 {
		_, ok := mp[v]
		if !ok {
			continue
		}
		mp[v] = [3]int{mp[v][0], i, 2}
	}

	var (
		result []string
		minVal = int(10e9)
	)

	for k, v := range mp {
		if v[2] == 1 {
			continue
		}
		if v[0]+v[1] == minVal {
			result = append(result, k)
			continue
		}

		if v[0]+v[1] < minVal {
			minVal = v[0] + v[1]
			result = nil
			result = append(result, k)
		}
		// fmt.Println(result, minVal)
	}

	return result
}
