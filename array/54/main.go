package main

import "fmt"

func main() {
	//fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

func spiralOrder(matrix [][]int) []int {
	var result []int

	left, right := 0, len(matrix[0])-1
	bot, top := 0, len(matrix)-1

	for left <= right && bot <= top {
		//fmt.Println(left, right, bot, top)
		for i := left; i <= right; i++ {
			result = append(result, matrix[bot][i])
			//fmt.Println(matrix[bot][i], "top side")
		}
		for i := bot + 1; i <= top; i++ {
			result = append(result, matrix[i][right])
			//fmt.Println(matrix[i][right], "right side")
		}
		if bot != top {
			for i := right - 1; i >= left; i-- {
				result = append(result, matrix[top][i])
				//fmt.Println(matrix[top][i], "bot side")
			}
		}

		if right != left {
			for i := top - 1; i >= bot+1; i-- {
				result = append(result, matrix[i][left])
				//fmt.Println(matrix[i][left], "left side")
			}
		}

		left++
		right--
		bot++
		top--
		//fmt.Println()
	}
	//fmt.Println()
	return result
}
