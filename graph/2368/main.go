package main

func main() {

}

func reachableNodes(n int, edges [][]int, restricted []int) int {
	graph := make([][]int, n)

	for _, v := range edges {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	res := make(map[int]bool)
	for _, v := range restricted {
		res[v] = true
	}

	var result int

	dfs(graph, 0, &result, make(map[int]bool), res)

	return result
}

func dfs(graph [][]int, curr int, result *int, visited, res map[int]bool) {
	if visited[curr] || res[curr] {
		return
	}
	visited[curr] = true

	for _, v := range graph[curr] {
		dfs(graph, v, result, visited, res)
	}

	*result++
}
