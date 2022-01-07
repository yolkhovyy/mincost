package main

import (
	"fmt"
	"math"
)

// ? https://www.geeksforgeeks.org/min-cost-path-dp-6/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minCost(cost *[][]int, m int, n int) int {
	result := 0
	if n < 0 || m < 0 {
		result = math.MaxInt64
	} else if m == 0 && n == 0 {
		result = (*cost)[m][n]
	} else {
		result = (*cost)[m][n] +
			min(minCost(cost, m-1, n),
				minCost(cost, m, n-1))
	}
	return result
}

func main() {
	tiles := [][]int{
		{5, 0, 8, 3},
		{0, 2, 1, 3},
		{1, 0, 9, 4},
		{0, 1, 0, 5}}

	fmt.Printf("%d", minCost(&tiles, 3, 3))
}
