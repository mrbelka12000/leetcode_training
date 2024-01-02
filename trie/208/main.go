package main

import "fmt"

func main() {
	t := Constructor()

	t.Insert("abc")
	t.Insert("check")

	fmt.Println(t.StartsWith("ab"))

	fmt.Println(t.Search("checks"))
}

type Trie struct {
	n *node
}

type node struct {
	end   bool
	child map[byte]*node
}

func newNode() *node {
	return &node{
		child: make(map[byte]*node),
	}
}

func Constructor() Trie {
	return Trie{
		n: newNode(),
	}
}

func (this *Trie) Insert(word string) {
	curr := this.n

	for i := 0; i < len(word); i++ {
		_, ok := curr.child[word[i]]
		if !ok {
			curr.child[word[i]] = newNode()
		}

		curr = curr.child[word[i]]
	}

	curr.end = true
}

func (this *Trie) Search(word string) bool {
	curr := this.n

	for i := 0; i < len(word); i++ {
		if _, ok := curr.child[word[i]]; !ok {
			return false
		}
		curr = curr.child[word[i]]
	}
	return curr.end
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.n

	for i := 0; i < len(prefix); i++ {
		if _, ok := curr.child[prefix[i]]; !ok {
			return false
		}
		curr = curr.child[prefix[i]]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
