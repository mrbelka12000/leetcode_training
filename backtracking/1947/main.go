package main

func main() {

}

func maxCompatibilitySum(st [][]int, mt [][]int) int {
	var result int

	visited = make(map[int]bool)
	used = make(map[int]bool)

	runner(st, mt, nil, 0, &result)

	return result
}

var (
	visited map[int]bool
	used    map[int]bool
)

func runner(st, mt [][]int, currSt []int, tmp int, result *int) {
	var delInd int
	if len(currSt) != 0 {
		score, i := getComparable(mt, currSt)
		delInd = i
		tmp += score
		// fmt.Println(used, visited, currSt, tmp)
	}

	if len(visited) == len(st) {
		*result = max(*result, tmp)
	}

	for ind, s := range st {
		if visited[ind] {
			continue
		}

		visited[ind] = true

		runner(st, mt, s, tmp, result)
		delete(visited, ind)
	}

	delete(used, delInd)
}

func getComparable(mt [][]int, st []int) (int, int) {
	var (
		maxVal, maxInd = -1, -1
	)

	for ind, m := range mt {
		if used[ind] {
			continue
		}
		var count int
		for j, v := range m {
			if st[j] == v {
				count++
			}
		}

		if maxVal < count {
			maxVal = count
			maxInd = ind
		}
	}

	used[maxInd] = true
	return maxVal, maxInd
}
