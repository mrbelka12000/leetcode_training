package main

import (
	"fmt"
)

func main() {
	fmt.Println(finalPositionOfSnake(3, []string{"DOWN", "RIGHT", "UP"}))
}

func finalPositionOfSnake(n int, commands []string) int {

	var ind int
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)

		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = ind
			ind++
		}
	}

	x, y := 0, 0
	for _, command := range commands {

		switch command {
		case "UP":
			x--
		case "DOWN":
			x++
		case "LEFT":
			y--
		case "RIGHT":
			y++
		}
	}

	return matrix[x][y]
}
