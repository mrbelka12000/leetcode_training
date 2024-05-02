package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}))

}

func removeSubfolders(folder []string) []string {
	t := &Trie{
		n: &node{
			child: make(map[string]*node),
		},
	}

	for _, v := range folder {
		t.Insert(v)
	}

	var result []string

	runner(t.n, "", &result)

	return result
}

type Trie struct {
	n *node
}

type node struct {
	isDir bool
	child map[string]*node
}

func newNode() *node {
	return &node{
		child: make(map[string]*node),
	}
}

func (t *Trie) Insert(word string) {
	curr := t.n
	words := strings.Split(word, "/")
	// fmt.Println(words)
	for i := 0; i < len(words); i++ {
		if words[i] == "" {
			continue
		}

		_, ok := curr.child[words[i]]
		if !ok {
			curr.child[words[i]] = newNode()
			// fmt.Println(string(words[i]), "suka")
		}

		curr = curr.child[words[i]]
		if curr.isDir {
			return
		}
	}

	curr.isDir = true
}

func runner(n *node, tmp string, result *[]string) {
	// fmt.Printf("%+v   %s\n", n, tmp)
	for k, v := range n.child {
		if v.isDir {
			*result = append(*result, tmp+"/"+string(k))
			continue
		}

		runner(v, tmp+"/"+k, result)
	}
}
