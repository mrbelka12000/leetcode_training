package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func countComponents(n int, edges [][]int) int {
	u := &uf{
		arr: make([]int, n),
	}

	for i := 0; i < len(u.arr); i++ {
		u.arr[i] = i
	}

	for _, v := range edges {
		u.union(v[0], v[1])
	}

	mp := make(map[int]bool)

	for _, v := range u.arr {
		mp[u.find(v)] = true
	}

	return len(mp)
}

type uf struct {
	arr []int
}

func (u *uf) find(x int) int {
	if u.arr[x] == x {
		return u.arr[x]
	}

	return u.find(u.arr[x])
}

func (u *uf) union(x, y int) {
	xn, yn := u.find(x), u.find(y)
	u.arr[yn] = xn
}
