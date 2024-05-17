package main

import (
	"container/list"
	"fmt"
)

func main() {

}

type (
	LFUCache struct {
		store     map[int]*list.Element
		storeList map[int]*list.List
		capacity  int
		min       int
	}

	Element struct {
		Key   int
		Value int
		Count int
	}
)

func Constructor(capacity int) LFUCache {
	return LFUCache{
		store:     make(map[int]*list.Element),
		storeList: make(map[int]*list.List),
		capacity:  capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	// defer func(){
	//     fmt.Println(this.store, "GET", key, this.storeList, this.min)
	// }()

	val, ok := this.store[key]
	if !ok {
		return -1
	}

	el := val.Value.(*Element)

	l := this.storeList[el.Count]
	l.Remove(val)

	el.Count++
	newList, ok := this.storeList[el.Count]
	if !ok {
		newList = list.New()
	}

	newVal := newList.PushBack(el)
	this.storeList[el.Count] = newList

	this.store[key] = newVal

	if l.Len() == 0 && this.min == el.Count-1 {
		this.min++
	}

	return val.Value.(*Element).Value
}

func (this *LFUCache) Put(key int, value int) {
	// defer func(){
	//     fmt.Println(this.store, "PUT", key, value, this.storeList, this.min)
	// }()

	if val, ok := this.store[key]; ok {
		val.Value.(*Element).Value = value
		this.Get(key)
		return
	}

	if len(this.store) == this.capacity {
		l := this.storeList[this.min]
		least := l.Front()
		l.Remove(least)
		delete(this.store, least.Value.(*Element).Key)
	}

	n := &Element{
		Key:   key,
		Value: value,
		Count: 1,
	}

	this.min = 1
	l, ok := this.storeList[1]
	if !ok {
		l = list.New()
	}

	el := l.PushBack(n)
	this.storeList[1] = l

	this.store[key] = el
}

func (this *LFUCache) check(key int, elem *list.Element) {
	el := elem.Value.(*Element)
	el.Count++
	this.store[key] = elem
}

func Print(l *list.List) {
	fmt.Println("\nPRINTER")

	elem := l.Front()
	for elem != nil {
		fmt.Printf("%+v\n", elem.Value.(*Element))
		elem = elem.Next()
	}

	fmt.Println("DONE\n")
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
