package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	//fmt.Println(spiralOrder([][]int{{1, 2, 3}}))
	//fmt.Println(spiralOrder([][]int{{1}, {2}}))
}

func spiralOrder(matrix [][]int) []int {
	var result []int
	top := len(matrix[0]) - 1
	right := len(matrix) - 1
	bot := 0
	left := 0
	var x, y, i int

	for bot <= top && left <= right {
		x, y = i, i
		fmt.Printf("top=%v, bot=%v, right=%v, left=%v, x=%v, y=%v\n", top, bot, right, left, x, y)

		for y < top {
			result = append(result, matrix[x][y])
			y++
		}
		fmt.Println(result, x, y, "first")

		for x < right {
			result = append(result, matrix[x][y])
			x++
		}
		fmt.Println(result, x, y, "second")

		for y > bot {
			result = append(result, matrix[x][y])
			y--
		}
		fmt.Println(result, x, y, "third")

		for x > left {
			result = append(result, matrix[x][y])
			x--
		}
		fmt.Println(result, x, y, "fourth")

		bot++
		top--
		left++
		right--
		i++
	}

	return result
}
