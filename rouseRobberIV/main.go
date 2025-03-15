func canSteal(nums []int, k, cap int) bool {
	count := 0
	i := 0
	for i < len(nums) {
		if nums[i] <= cap {
			count++
			i += 2
		} else {
			i++
		}
	}
	return count >= k
}

func minCapability(nums []int, k int) int {
    low, high := math.MaxInt32, math.MinInt32

	for _, num := range nums {
		if num < low {
			low = num
		}
		if num > high {
			high = num
		}
	}

	for low < high {
		mid := low + (high-low)/2
		if canSteal(nums, k, mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

