package main

func main() {

}

func solve(board [][]byte) {
	var (
		dirs  = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		dfs   func(x, y int)
		isSur bool
		path  [][2]int
	)

	visited := make(map[[2]int]bool)

	dfs = func(x, y int) {
		path = append(path, [2]int{x, y})
		if x == 0 || x == len(board)-1 || y == 0 || y == len(board[0])-1 {
			isSur = true
			return
		}

		visited[[2]int{x, y}] = true

		for _, dir := range dirs {
			xx, yy := x+dir[0], y+dir[1]
			if xx < 0 || xx >= len(board) || yy < 0 || yy >= len(board[0]) || board[xx][yy] == 'X' || visited[[2]int{xx, yy}] {
				continue
			}

			// fmt.Println(xx,yy)
			dfs(xx, yy)
		}

		return
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				clear(visited)
				path = make([][2]int, 0)
				isSur = false
				dfs(i, j)

				if !isSur {
					for _, v := range path {
						board[v[0]][v[1]] = 'X'
					}
				}
			}
		}
	}
}
