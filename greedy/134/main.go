package main

import (
	"fmt"
)

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}

func canCompleteCircuit(gas []int, cost []int) int {
	var inds []int

	for i := 0; i < len(gas); i++ {
		if gas[i] >= cost[i] {
			inds = append(inds, i)
		}
	}

	for len(inds) != 0 {
		var (
			ind = inds[0]
			tmp = gas[ind] - cost[ind]
			i   int
		)

		inds = inds[1:]

		for i = ind + 1; i != ind; i++ {
			if i == len(gas) {
				i = -1
				continue
			}
			if tmp+gas[i] < cost[i] {
				break
			}

			if len(inds) > 0 && inds[0] == i {
				inds = inds[1:]
			}

			tmp += gas[i] - cost[i]
		}

		if i == ind {
			return ind
		}
	}

	return -1
}
