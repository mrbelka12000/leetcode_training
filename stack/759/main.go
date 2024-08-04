package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello world")
}

type Interval struct {
	Start int
	End   int
}

func employeeFreeTime(schedule [][]*Interval) []*Interval {
	var it []*Interval
	for _, v := range schedule {
		for _, i := range v {
			it = append(it, i)
		}
	}
	sort.Slice(it, func(i, j int) bool {
		return it[i].Start < it[j].Start
	})

	stack := []*Interval{it[0]}

	for i := 1; i < len(it); i++ {
		if stack[len(stack)-1].End >= it[i].Start {
			if stack[len(stack)-1].End < it[i].End {
				stack[len(stack)-1].End = it[i].End
			}
			continue
		}
		stack = append(stack, it[i])
	}

	var result []*Interval

	for i := 0; i < len(stack)-1; i++ {
		if stack[i+1].Start-stack[i].End > 0 {
			result = append(result, &Interval{
				Start: stack[i].End,
				End:   stack[i+1].Start,
			})
		}
	}

	return result
}
