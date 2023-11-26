package main

import "fmt"

func main() {

	wd := Constructor()

	wd.AddWord("add")

	fmt.Println(wd.Search("a.d."))
}

type WordDictionary struct {
	n *node
}

type node struct {
	child map[byte]*node
	char  string
	end   bool
}

func newNode() *node {
	return &node{
		child: make(map[byte]*node),
	}
}

func Constructor() WordDictionary {
	return WordDictionary{
		n: newNode(),
	}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.n

	for i := 0; i < len(word); i++ {
		if _, ok := curr.child[word[i]]; !ok {
			curr.child[word[i]] = newNode()
			curr.child[word[i]].char = string(word[i])
		}
		curr = curr.child[word[i]]
	}

	curr.end = true
}

func (this *WordDictionary) Search(word string) bool {
	curr := this.n

	return search(word, 0, curr)
}

func search(word string, ind int, n *node) bool {

	var notFound bool
	for i := ind; i < len(word); i++ {
		// fmt.Println(ind, i, n, string(word[i]))
		if word[i] == '.' {
			if len(n.child) == 0 {
				return false
			}
			for _, v := range n.child {
				if search(word, i+1, v) {
					return true
				} else {
					notFound = true
				}
			}
			break
		}

		//for k, v := range n.child {
		//	fmt.Println(string(k), v.end)
		//}
		//
		//fmt.Println()
		var ok bool
		// fmt.Println(string(word[i]), n, "1324123")
		_, ok = n.child[word[i]]
		if !ok {
			return false
		}
		n = n.child[word[i]]
	}

	if notFound {
		return false
	}
	return n.end
}
