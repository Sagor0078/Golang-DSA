package main

import (
	"fmt"
)

func digitSum(x int) int {
	sum := 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	return sum
}

func maximumSum(nums []int) int {
	maxD := make([]int, 82) // Array to store maximum number for each digit sum
	ans := -1

	for _, x := range nums {
		D := digitSum(x)
		if maxD[D] > 0 {
			ans = max(ans, maxD[D]+x)
		}
		maxD[D] = max(maxD[D], x)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
