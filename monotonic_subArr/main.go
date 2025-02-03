package main

func longestMonotonicSubarray(nums []int) int {
	ans, increment, decrement := 0, 1, 1
	if len(nums) == 1 {
		return 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			increment++
			decrement = 1
		} else if nums[i-1] > nums[i] {
			decrement++
			increment = 1
		} else {
			increment = 1
			decrement = 1
		}
		ans = max(ans, increment)
		ans = max(ans, decrement)
	}
	return ans
}
