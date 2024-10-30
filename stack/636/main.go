package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(exclusiveTime(1, []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}))
	fmt.Println(exclusiveTime(2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}))
}

func exclusiveTime(n int, logs []string) []int {
	result := make([]int, n)
	var stack []info

	for _, v := range logs {
		id, val, isStart := parseLog(v)
		if isStart {
			stack = append(stack, info{
				id:  id,
				val: val,
			})
		} else {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			diff := val - last.val + 1
			result[last.id] += diff
			if len(stack) > 0 {
				result[stack[len(stack)-1].id] -= diff
			}

		}
	}
	return result
}

func parseLog(log string) (int, int, bool) {
	arr := strings.Split(log, ":")
	if len(arr) != 3 {
		panic("suka")
	}

	id, _ := strconv.Atoi(arr[0])
	isStart := arr[1] == "start"
	val, _ := strconv.Atoi(arr[2])

	return id, val, isStart
}

type info struct {
	id  int
	val int
}
