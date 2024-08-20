package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type ImmutableListNode struct {
}

func (this *ImmutableListNode) getNext() ImmutableListNode {
	return ImmutableListNode{}
}

func (this *ImmutableListNode) printValue() {
	// print the value of this node.
}

/*
n = 1000
[-1000:1000]
*/

func printLinkedListInReverse(head ImmutableListNode) {
	// Leetcode issue
	//if head == nil {
	//	return
	//}

	printLinkedListInReverse(head.getNext())

	head.printValue()
}
