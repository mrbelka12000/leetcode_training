package main

func main() {

}

func findChampion(n int, edges [][]int) int {
	arr := make([]int, n)
	for _, v := range edges {
		arr[v[1]]++
	}

	var result int = -1

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			if result != -1 {
				return -1
			}
			result = i
		}
	}

	return result
}
