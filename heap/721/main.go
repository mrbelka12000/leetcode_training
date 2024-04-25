package main

import (
	"sort"
)

func main() {

}

func accountsMerge(accounts [][]string) [][]string {
	mp := make(map[string][]int)
	for ind, acc := range accounts {
		for i := 1; i < len(acc); i++ {
			mp[acc[i]] = append(mp[acc[i]], ind)
		}
	}

	visited = make(map[string]bool)
	var result [][]string
	for len(mp) != 0 {
		var (
			maxEm  string
			maxVal int
			name   string
		)

		for k, v := range mp {
			if visited[k] {
				delete(mp, k)
				continue
			}

			if maxVal < len(v) {
				maxVal = len(v)
				name = accounts[v[0]][0]
				maxEm = k
			}
		}

		if len(mp) == 0 {
			break
		}

		emails := []string{""}

		dfs(mp, accounts, maxEm, &emails)

		sort.Slice(emails, func(i, j int) bool {
			return emails[i] < emails[j]
		})

		emails[0] = name

		result = append(result, emails)
	}

	return result
}

var visited map[string]bool

func dfs(graph map[string][]int, accounts [][]string, curr string, tmp *[]string) {

	for _, v := range graph[curr] {
		for i, acc := range accounts[v] {
			if i == 0 {
				continue
			}
			if visited[acc] {
				continue
			}
			visited[acc] = true

			*tmp = append(*tmp, acc)
			dfs(graph, accounts, acc, tmp)
		}
	}
}
