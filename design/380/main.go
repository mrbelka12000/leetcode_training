package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rs := Constructor()

	rs.Insert(1)
	rs.Insert(2)
	rs.Insert(3)

	rs.Remove(3)

	fmt.Println(rs.GetRandom())
	fmt.Println(rs.GetRandom())
}

type RandomizedSet struct {
	arr []int
	mp  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		mp: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.mp[val]

	if !ok {
		this.arr = append(this.arr, val)
	}
	this.mp[val] = len(this.arr) - 1

	return !ok
}

func (this *RandomizedSet) Remove(val int) bool {
	ind, ok := this.mp[val]

	if ok {
		this.remove(this.arr, ind)
	}

	delete(this.mp, val)

	return ok
}

func (this *RandomizedSet) GetRandom() int {
	val := this.arr[rand.Intn(len(this.arr))]
	for val == -149437329475 {
		val = this.arr[rand.Intn(len(this.arr))]
	}
	return val
}

func (this *RandomizedSet) remove(s []int, i int) {
	this.arr[i] = -149437329475
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
