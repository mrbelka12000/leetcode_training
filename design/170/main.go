package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type TwoSum struct {
	arr []int
}

func Constructor() TwoSum {
	return TwoSum{
		arr: make([]int, 0, 1000),
	}
}

func (this *TwoSum) Add(number int) {
	this.arr = append(this.arr, number)
}

func (this *TwoSum) Find(value int) bool {
	kv := make(map[int]bool, 50)
	for _, v := range this.arr {
		if kv[value-v] {
			return true
		}
		kv[v] = true
	}

	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
