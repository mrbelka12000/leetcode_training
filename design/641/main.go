package main

import "container/list"

func main() {

}

type (
	MyCircularDeque struct {
		l *list.List
		k int
	}
	Element struct {
		Value int
	}
)

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		l: list.New(),
		k: k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.l.Len() == this.k {
		return false
	}
	this.l.PushFront(&Element{Value: value})
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.l.Len() == this.k {
		return false
	}
	this.l.PushBack(&Element{Value: value})
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.l.Len() == 0 {
		return false
	}

	this.l.Remove(this.l.Front())
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.l.Len() == 0 {
		return false
	}
	this.l.Remove(this.l.Back())
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.l.Len() == 0 {
		return -1
	}

	return this.l.Front().Value.(*Element).Value
}

func (this *MyCircularDeque) GetRear() int {
	if this.l.Len() == 0 {
		return -1
	}

	return this.l.Back().Value.(*Element).Value
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.l.Len() == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.l.Len() == this.k
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
