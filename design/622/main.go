package main

func main() {
}

type MyCircularQueue struct {
	k    int
	list []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		k: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.list) == this.k {
		return false
	}

	this.list = append(this.list, value)

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if len(this.list) == 0 {
		return false
	}

	this.list = this.list[1:]
	return true
}

func (this *MyCircularQueue) Front() int {
	if len(this.list) == 0 {
		return -1
	}

	return this.list[0]
}

func (this *MyCircularQueue) Rear() int {
	if len(this.list) == 0 {
		return -1
	}

	return this.list[len(this.list)-1]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.list) == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return len(this.list) == this.k
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
