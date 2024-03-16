package main

import "fmt"

func main() {
	fmt.Println(findMaxLength([]int{0, 1, 0, 1}))

}

func findMaxLength(nums []int) int {
	mp := make(map[int]int)
	mp[0] = -1
	var count, result int

	for i, v := range nums {
		switch v {
		case 0:
			count--
		case 1:
			count++
		}

		if val, ok := mp[count]; ok {
			result = max(result, i-val)
		} else {
			mp[count] = i
		}
	}

	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
