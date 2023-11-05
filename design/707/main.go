package main

import "fmt"

func main() {
	ll := Constructor()
	ll.AddAtHead(7)
	ll.AddAtHead(2)
	ll.AddAtHead(1)
	ll.AddAtIndex(3, 0)
	ll.DeleteAtIndex(2)
	ll.AddAtHead(6)
	ll.AddAtTail(1)
	//ll.AddAtTail(2)
	//ll.AddAtIndex(0, 1)
	//ll.AddAtIndex(2, 3)
	ll.print()
	fmt.Println(ll.Get(1))
	//ll.DeleteAtIndex(1)

	//ll.print()
}

type MyLinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value int
	Next  *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	var i int
	curr := this.Head
	for curr != nil {
		if i == index {
			return curr.Value
		}
		i++
		curr = curr.Next
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	tmp := &Node{
		Value: val,
	}
	if this.Head == nil {
		this.Head = tmp
	} else {
		tmp.Next = this.Head
		this.Head = tmp
	}
	if this.Tail == nil {
		this.Tail = tmp
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.Head == nil {
		this.AddAtHead(val)
		return
	}

	curr := this.Head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = &Node{Value: val}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	tmp := &Node{
		Value: val,
	}
	if index == 0 {
		tmp.Next = this.Head
		this.Head = tmp
		return
	}

	var i int
	curr := this.Head
	for curr != nil {
		i++
		if index == i {
			//if curr.Next != nil {
			tmp.Next = curr.Next
			//}
			curr.Next = tmp
		}
		//fmt.Println(curr.Value, tmp.Value, "value")
		curr = curr.Next
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.Head = this.Head.Next
		return
	}

	var i int
	curr := this.Head
	for curr != nil && curr.Next != nil {
		i++
		if i == index {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}
}

func (this MyLinkedList) print() {
	for this.Head != nil {
		fmt.Println(this.Head.Value)
		this.Head = this.Head.Next
	}
	fmt.Println("print")
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
