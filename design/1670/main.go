package main

import "fmt"

func main() {
	q := Constructor()
	q.PushBack(1)
	q.PushBack(2)
	q.PushBack(3)
	q.PushBack(4)
	q.PushBack(5)
	q.PushMiddle(6)
	fmt.Println(q.q)

	fmt.Println(q.PopMiddle())
}

type FrontMiddleBackQueue struct {
	q []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.q = append([]int{val}, this.q...)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	half := make([]int, len(this.q)/2)
	copy(half, this.q[:len(this.q)/2])

	half = append(half, val)
	this.q = append(half, this.q[len(this.q)/2:]...)
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.q = append(this.q, val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if len(this.q) == 0 {
		return -1
	}
	val := this.q[0]
	this.q = this.q[1:]
	return val
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if len(this.q) == 0 {
		return -1
	}
	var ind int

	if len(this.q)%2 == 0 {
		ind = len(this.q)/2 - 1
	} else {
		ind = len(this.q) / 2
	}

	val := this.q[ind]
	this.q = remove(this.q, ind)
	return val
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if len(this.q) == 0 {
		return -1
	}
	val := this.q[len(this.q)-1]
	this.q = this.q[:len(this.q)-1]
	return val
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
