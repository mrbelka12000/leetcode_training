package main

import (
	"fmt"
)

func main() {
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
}

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return amount
	}

	// Init Steps
	steps := make(map[int]int, len(coins))
	for _, coin := range coins {
		if coin == amount {
			return 1
		}

		steps[coin] = 1
	}

	var (
		used   = make(map[int]bool, len(coins))
		queue  = append([]int{}, coins...)
		depth  = 0
		result = -1
	)

	for len(queue) > 0 && result == -1 {
		size := len(queue)
		// fmt.Println("result", result, queue)
		for i := 0; i < size; i++ {
			sum := queue[0]
			queue = queue[1:]
			// fmt.Println(cur)
			if sum == amount {
				return depth
			}

			for _, coin := range coins {
				newSum := sum + coin
				if newSum > amount {
					continue
				}
				if newSum == amount {
					return depth + 2
				}
				if used[newSum] {
					continue
				}
				used[newSum] = true

				remain := amount - newSum

				// fmt.Println("sum", sum, "coin", coin, "newSum", newSum, "remain", remain, steps)
				if step, ok := steps[remain]; ok {
					t := step + 2 + depth
					if result == -1 || result > t {
						result = t
					}
					// fmt.Println("result:", result)
				} else {
					steps[remain] = depth + 2
				}

				queue = append(queue, newSum)
			}
		}
		depth++
	}

	return result
}
