package main

import "fmt"

func main() {
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2))
}

func findPairs(nums []int, k int) int {
	var result int

	mp := make(map[[2]int]struct{})
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if abs(nums[i]-nums[j]) == k {
				key := getKeys(nums[i], nums[j])
				if _, ok := mp[key[0]]; ok {
					continue
				}
				if _, ok := mp[key[1]]; ok {
					continue
				}

				mp[key[0]] = struct{}{}
				mp[key[1]] = struct{}{}

				result++
			}
		}
	}
	return result
}

func getKeys(a, b int) (key [][2]int) {
	key = append(key, [2]int{a, b})
	key = append(key, [2]int{b, a})
	return
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
