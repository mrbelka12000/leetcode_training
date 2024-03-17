package main

func main() {

}

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	return false
}
func (this *Iterator) next() int {
	return 0
}

type PeekingIterator struct {
	iter *Iterator
	hp   bool
	elem *int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter: iter,
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.hp || this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if !this.hp {
		return this.iter.next()
	}

	this.hp = false
	result := *this.elem
	this.elem = nil

	return result
}

func (this *PeekingIterator) peek() int {
	if !this.hp {
		this.hp = true
		val := this.iter.next()
		this.elem = &val
	}

	return *this.elem
}
