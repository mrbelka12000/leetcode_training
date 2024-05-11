package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {

	res := strings.Builder{}

	for head != nil {
		res.WriteString(fmt.Sprint(head.Val))
		head = head.Next
	}

	if i, err := strconv.ParseInt(res.String(), 2, 64); err != nil {
		fmt.Println(err)
	} else {
		return int(i)
	}

	return 0
}
