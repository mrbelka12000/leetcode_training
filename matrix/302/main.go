package main

import (
	"fmt"
)

func main() {
	fmt.Println(minArea([][]byte{{'0', '0', '1', '0'}, {'0', '1', '1', '0'}, {'0', '1', '0', '0'}}, 0, 2))
}

func minArea(image [][]byte, x int, y int) int {
	coords := getCoords(image, x, y)

	var (
		xU int = len(image)
		xD int
		yR int = len(image[0])
		yL int
	)

	for _, v := range coords {
		xU = min(xU, v[0])
		xD = max(xD, v[0])

		yR = min(yR, v[1])
		yL = max(yL, v[1])
	}

	return (xD - xU + 1) * (yL - yR + 1)
}

func getCoords(image [][]byte, x, y int) [][]int {
	if x < 0 || x >= len(image) || y < 0 || y >= len(image[0]) || image[x][y] == '0' {
		return nil
	}

	image[x][y] = '0'

	arr := [][]int{{x, y}}

	arr = append(arr, getCoords(image, x+1, y)...)
	arr = append(arr, getCoords(image, x-1, y)...)
	arr = append(arr, getCoords(image, x, y+1)...)
	arr = append(arr, getCoords(image, x, y-1)...)

	return arr
}
