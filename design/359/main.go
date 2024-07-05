package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

type Logger struct {
	store map[string]int
}

func Constructor() Logger {
	return Logger{
		store: make(map[string]int, 1000),
	}
}

func (this *Logger) ShouldPrintMessage(t int, m string) bool {
	if val, ok := this.store[m]; ok {
		if val > t {
			return false
		}
	}
	this.store[m] = t + 10
	return true
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
