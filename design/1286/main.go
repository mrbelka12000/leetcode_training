package main

import (
	"fmt"
)

func main() {
	c := Constructor("abc", 2)

	fmt.Println(c.Next())
	fmt.Println(c.Next())
	fmt.Println(c.HasNext())
	fmt.Println(c.Next())
	fmt.Println(c.HasNext())

	//for i := 1; i <= 15; i++ {
	//	fmt.Println(c.Next())
	//}
	//for i := 1; i <= 29; i++ {
	//	fmt.Println(c.HasNext())
	//}
}

type CombinationIterator struct {
	c       string
	l       int
	start   int
	hasNext bool
	ch      chan string
	q       []string
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	c := CombinationIterator{
		c:       characters,
		l:       combinationLength,
		ch:      make(chan string, 5),
		hasNext: true,
	}

	go func() {
		c.helper(c.start, "")
		close(c.ch)
	}()

	for v := range c.ch {
		c.q = append(c.q, v)
	}

	return c
}

func (this *CombinationIterator) Next() string {
	val := this.q[0]
	this.q = this.q[1:]
	return val
}

func (this *CombinationIterator) helper(ind int, str string) {
	if len(str) == this.l {
		this.ch <- str
		return
	}

	for i := ind; i < len(this.c); i++ {
		this.helper(i+1, str+string(this.c[i]))
	}
}

func (this *CombinationIterator) HasNext() bool {
	return len(this.q) != 0
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
