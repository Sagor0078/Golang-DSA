package main

func maxAscendingSum(nums []int) int {
	ans, curr := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			curr += nums[i]
		} else {
			curr = nums[i]
		}
		ans = max(ans, curr)
	}
	return ans
}
