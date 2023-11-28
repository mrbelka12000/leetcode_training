package main

func main() {

}

type MapSum struct {
	root *node
	dic  map[string]int
}

type node struct {
	child map[string]*node
	end   bool
	value int
}

func newNode() *node {
	return &node{
		child: make(map[string]*node),
	}
}

func Constructor() MapSum {
	return MapSum{
		root: newNode(),
		dic:  make(map[string]int),
	}
}

func (this *MapSum) Insert(key string, val int) {
	value, exists := this.dic[key]
	curr := this.root
	for i := 0; i < len(key); i++ {
		if _, ok := curr.child[string(key[0:i+1])]; !ok {
			curr.child[string(key[0:i+1])] = newNode()
		}

		curr = curr.child[string(key[0:i+1])]
		// fmt.Println(string(key[0:i+1]),exists,curr,value,val)
		if !exists {
			curr.value += val
		} else {
			curr.value = curr.value - value + val
		}
	}
	// fmt.Println()

	this.dic[key] = val
	curr.end = true
}

func (this *MapSum) Sum(prefix string) int {
	curr := this.root

	for i := 0; i < len(prefix); i++ {
		if _, ok := curr.child[string(prefix[0:i+1])]; !ok {
			return 0
		}
		curr = curr.child[string(prefix[0:i+1])]
	}

	return curr.value
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
