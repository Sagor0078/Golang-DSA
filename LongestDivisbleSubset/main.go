import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	ans := []int{}

	if n == 0 {
		return ans
	}

	sort.Ints(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	maxVal := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				if dp[i] > maxVal {
					maxVal = dp[i]
				}
			}
		}
	}

	prev := -1
	for i := n - 1; i >= 0; i-- {
		if dp[i] == maxVal && (prev == -1 || prev%nums[i] == 0) {
			ans = append(ans, nums[i])
			prev = nums[i]
			maxVal--
		}
	}

	return ans
}
