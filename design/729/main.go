package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello world")
}

type MyCalendar struct {
	store [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	// fmt.Println(this.store, start , end, count)
	for _, v := range this.store {
		if v[1] > start && v[1] < end {
			return false
		}

		if start < v[0] && end <= v[1] && end > v[0] {
			return false
		}

		if start >= v[0] && start < v[1] {
			return false
		}
	}
	// fmt.Println("ok")
	this.store = append(this.store, []int{start, end})

	sort.Slice(this.store, func(i, j int) bool {
		return this.store[i][0] < this.store[j][0]
	})

	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
