package main

import (
	"fmt"
)

func main() {
	fmt.Println(findAllRecipes([]string{"bread"}, [][]string{{"yeast", "flour"}}, []string{"yeast", "flour", "corn"}))
}

func findAllRecipes(rec []string, ing [][]string, sup []string) []string {
	graph := make(map[string][]string)
	inDegree := make(map[string]int)
	for i, v := range ing {
		for _, in := range v {
			graph[in] = append(graph[in], rec[i])
		}

		inDegree[rec[i]] = len(v)
	}

	result := make([]string, 0, len(rec))
	for len(sup) != 0 {
		n := sup[0]
		sup = sup[1:]
		for _, v := range graph[n] {
			inDegree[v]--
			if inDegree[v] == 0 {
				sup = append(sup, v)
				result = append(result, v)
			}
		}
	}

	return result
}
