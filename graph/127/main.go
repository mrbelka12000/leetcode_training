package main

import "fmt"

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	g := &G{
		store: make(map[string]*node),
	}
	var check bool
	for i := 0; i < len(wordList); i++ {
		if wordList[i] == endWord {
			check = true
		}
		if isEdje(beginWord, wordList[i]) {
			g.addAdj(beginWord, wordList[i])
			g.addAdj(wordList[i], beginWord)
		}

		for j := i + 1; j < len(wordList); j++ {
			if isEdje(wordList[i], wordList[j]) {
				g.addAdj(wordList[i], wordList[j])
				g.addAdj(wordList[j], wordList[i])
			}
		}
	}

	if !check {
		return 0
	}

	result, ok := g.dfs(beginWord, endWord, make(map[string]bool), 0, 0)
	if !ok {
		return 0
	}

	return result + 1
}

func (g *G) dfs(curr, end string, visited map[string]bool, dep int, result int) (int, bool) {
	q := []*node{g.getNode(curr)}

	for len(q) != 0 {

		size := len(q)

		for i := 0; i < size; i++ {

			n := q[0]
			q = q[1:]
			if n.val == end {
				return result, true
			}
			visited[n.val] = true

			for _, v := range n.adj {
				if visited[v.val] {
					continue
				}
				q = append(q, v)
			}
		}
		result++
	}
	return result, false
}

func isEdje(s1, s2 string) bool {
	if s1 == s2 {
		return false
	}

	arr := []byte(s1)

	for i := 0; i < len(s2); i++ {
		if arr[i] == s2[i] {
			arr[i] = '0'
		}
	}

	var tmp int
	for i := 0; i < len(arr); i++ {
		if arr[i] != '0' {
			tmp++
		}
	}

	return tmp <= 1
}

type (
	G struct {
		store map[string]*node
	}

	node struct {
		val string
		adj map[string]*node
	}
)

func (g *G) addNode(val string) *node {
	n := &node{
		val: val,
		adj: make(map[string]*node),
	}

	g.store[val] = n
	return n
}

func (g *G) getNode(val string) *node {
	if val, ok := g.store[val]; ok {
		return val
	}
	return g.addNode(val)
}

func (g *G) addAdj(from, to string) {
	fromNode, toNode := g.getNode(from), g.getNode(to)
	for _, v := range fromNode.adj {
		if v.val == to {
			return
		}
	}
	fromNode.adj[to] = toNode
}

func (g *G) Print() {
	for k, v := range g.store {
		fmt.Printf("node %v has adj:\n", k)
		for _, val := range v.adj {
			fmt.Println(val.val)
		}
	}
}
