package main

import (
	"fmt"
)

func main() {
	l := Constructor(2)
	//[[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]

	l.Put(2, 1)
	l.Put(1, 1)
	l.Put(2, 3)
}

type LRUCache struct {
	store map[int]*Node
	head  *Node
	tail  *Node
	cp    int
}

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		store: make(map[int]*Node),
		cp:    capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	// defer func() {
	// 	Print(this.head)
	// }()

	node, ok := this.store[key]
	if !ok {
		return -1
	}

	// fmt.Printf("GET %+v\n", node)
	if node.Next == nil {
		return node.Val
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if this.head == node {
		this.head = this.head.Next
	}

	node.Next.Prev = node.Prev
	node.Next = nil
	node.Prev = this.tail
	this.tail.Next = node
	this.tail = node

	return node.Val
}

func Print(head *Node) {
	fmt.Println("PRINTER")
	for head != nil {
		fmt.Printf("%+v\n", head)
		head = head.Next
	}
	fmt.Println()
}

func (this *LRUCache) Put(key int, value int) {
	// fmt.Println(key, value, "PUT")

	// defer Print(this.head)

	var node *Node
	if v, ok := this.store[key]; !ok {
		if len(this.store) == this.cp {
			delete(this.store, this.head.Key)
			this.head = this.head.Next
		}
		node = &Node{
			Key: key,
			Val: value,
		}
		this.store[key] = node
	} else {
		v.Val = value
		if v.Next == nil {
			return
		}
		node = v
	}

	if this.head == nil {
		this.head = node
		this.tail = node
		return
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if this.head == node {
		this.head = this.head.Next
	}
	node.Next = nil
	node.Prev = this.tail
	this.tail.Next = node
	this.tail = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
